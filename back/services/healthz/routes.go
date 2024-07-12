package healthz

import (
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
	r.GET("/healthz", h.handleGetHealthz)
}

func (controllers *Handler) handleGetHealthz(ctx *gin.Context) {
	utils.RespondWithJSON(ctx, http.StatusOK, types.ApiResponse{
		Success: true,
		Data:    "Server is live",
	})
}
