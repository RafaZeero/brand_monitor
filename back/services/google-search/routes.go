package googlesearch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/RafaZeero/brand_monitor/types"
	"github.com/RafaZeero/brand_monitor/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/search", h.handleGetQuery)
}

func (controllers *Handler) handleGetQuery(ctx *gin.Context) {
	type SearchData struct {
		Items []string `json:"items"`
	}

	var searchData *SearchData
	if err := ctx.BindJSON(&searchData); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if len(searchData.Items) == 0 {
		utils.RespondWithError(ctx, http.StatusOK, "Empty query")
		return
	}

	data := make([]*types.CustomSearchResponse, 0)

	for _, item := range searchData.Items {
		gRes, err := SearchFor(item)
		if err != nil {
			utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		data = append(data, gRes)
	}
	utils.RespondWithJSON(ctx, http.StatusOK, types.ApiResponse{
		Success: true,
		Data:    data,
	})
}

func SearchFor(query string) (*types.CustomSearchResponse, error) {
	url := fmt.Sprintf(
		"https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s",
		utils.EnvOrFatal("GOOGLE_API_KEY"),
		utils.EnvOrFatal("GOOGLE_CX"),
		query,
	)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data types.CustomSearchResponse
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
