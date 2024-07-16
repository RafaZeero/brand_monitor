package healthz

import (
	"context"
	"net/http"
	"time"

	"github.com/RafaZeero/brand_monitor/types"
	"github.com/RafaZeero/brand_monitor/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	store types.TestStore
}

func NewHandler(store types.TestStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/healthz", h.handleGetHealthz)
	r.POST("/test", h.handleTesteAddData)
}

func (h *Handler) handleGetHealthz(ctx *gin.Context) {
	utils.RespondWithJSON(ctx, http.StatusOK, types.ApiResponse{
		Success: true,
		Data:    "Server is live",
	})
}

func (h *Handler) handleTesteAddData(ctx *gin.Context) {
	var data struct {
		Text string `json:"text"`
	}
	if err := ctx.BindJSON(&data); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.store.AddData(context.Background(), &types.TestAddData{
		ID:        primitive.NewObjectID(),
		Text:      data.Text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, types.ApiResponse{
		Success: true,
		Data:    "Data added successfully",
	})
}
