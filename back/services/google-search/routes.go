package googlesearch

import (
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
	r.POST("/search", h.handleCreateSearch)
	r.GET("/search", h.handleGetSearches)
}

func (h *Handler) handleCreateSearch(ctx *gin.Context) {
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
				competitors = append(competitors, i.DisplayLink)
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

func (h *Handler) handleGetSearches(ctx *gin.Context) {
	searches, err := h.store.GetSearches(ctx)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusOK, types.ApiResponse{
		Success: true,
		Data:    searches,
	})
}
