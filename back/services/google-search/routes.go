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

type Handler struct {
	store types.SearchStore
}

func NewHandler(store types.SearchStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/search", h.handleGetSearch)
}

func (h *Handler) handleGetSearch(ctx *gin.Context) {
	type SearchItems struct {
		Items []string `json:"items"`
	}

	var searchItems *SearchItems
	if err := ctx.BindJSON(&searchItems); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if len(searchItems.Items) == 0 {
		utils.RespondWithError(ctx, http.StatusOK, "Empty query")
		return
	}

	data := make([]*types.GoogleSearchApiResponse, 0)

	for _, item := range searchItems.Items {
		res, err := SearchFor(item)
		if err != nil {
			utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		data = append(data, res)

		competitors := func() []string {
			competitors := []string{}
			for _, i := range res.Items {
				competitors = append(competitors, i.Title)
			}
			return competitors
		}()

		if err := h.store.CreateSearch(ctx, &types.CreateSearchPayload{
			Term:        item,
			Competitors: competitors,
		}); err != nil {
			utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
	utils.RespondWithJSON(ctx, http.StatusOK, types.ApiResponse{
		Success: true,
		Data:    data,
	})
}

func SearchFor(query string) (*types.GoogleSearchApiResponse, error) {
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

	var data types.GoogleSearchApiResponse
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
