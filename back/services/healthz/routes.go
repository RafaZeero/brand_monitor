package healthz

import (
	"net/http"

	"github.com/RafaZeero/brand_monitor/types"
	"github.com/RafaZeero/brand_monitor/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store types.TestStore
}

func NewHandler(store types.TestStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/healthz", h.handleGetHealthz)
	r.GET("/ping", h.handleGetPing)
}

func (h *Handler) handleGetHealthz(ctx *gin.Context) {
	utils.RespondWithJSON(ctx, http.StatusOK, types.ApiResponse{
		Success: true,
		Data:    "Server is live",
	})
}

func (h *Handler) handleGetPing(ctx *gin.Context) {
	if err := h.store.Ping(); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, types.ApiResponse{
		Success: true,
		Data:    "Db connected",
	})
}
