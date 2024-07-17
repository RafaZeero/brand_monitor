package googlesearch

import (
	"net/http"

	"github.com/RafaZeero/brand_monitor/types"
	"github.com/RafaZeero/brand_monitor/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
		Items []string `json:"items" binding:"required"`
		Email string   `json:"email" binding:"required"`
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

	for _, item := range searchItems.Items {
		res, err := SearchFor(item)
		if err != nil {
			utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		competitors := func() []string {
			competitors := []string{}
			for _, i := range res.Items {
				competitors = append(competitors, i.DisplayLink)
			}
			return competitors
		}()

		if err := h.store.CreateSearch(ctx, &types.CreateSearchPayload{
			UserEmail:   searchItems.Email,
			Term:        item,
			Competitors: competitors,
		}); err != nil {
			utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
	utils.RespondWithJSON(ctx, http.StatusOK, types.ApiResponse{
		Success: true,
		Message: "Search created successfully!",
	})
}

func (h *Handler) handleGetSearches(ctx *gin.Context) {
	email := ctx.Query("email")
	if email == "" {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid email")
		return
	}

	searches, err := h.store.GetSearches(ctx, email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.RespondWithError(ctx, http.StatusOK, "No searches found for this email")
			return
		}
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusOK, types.ApiResponse{
		Success: true,
		Data:    searches,
	})
}
