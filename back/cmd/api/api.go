package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RafaZeero/brand_monitor/services/healthz"
	"github.com/RafaZeero/brand_monitor/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{addr: addr}
}

func (s *ApiServer) Run() error {
	// Zap logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Initialize Router
	r := gin.New()

	// Middlewares
	// CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
		MaxAge:          300,
	}))
	r.Use(gin.Logger())

	// Routes
	RegisterRoutes(r)

	// Server config
	srv := &http.Server{
		Addr:    ":" + s.addr,
		Handler: r,
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	logger.Info("[" + env + "]" + "Server running on port " + s.addr)

	// Run server
	return srv.ListenAndServe()
}

func RegisterRoutes(r *gin.Engine) {
	healthzHandler := healthz.NewHandler()

	v1 := r.Group("/v1")
	{
		healthzHandler.RegisterRoutes(v1)
		// others routes...
	}

	// Fallback para as rotas
	r.NoRoute(func(ctx *gin.Context) {
		// Ex: /abcdef
		url := ctx.Request.URL
		// Ex: /abcdef not-found
		utils.RespondWithJSON(ctx, http.StatusNotFound, struct {
			Success bool   `json:"success"`
			URL     string `json:"url"`
		}{
			Success: false,
			URL:     fmt.Sprintf("%v - url não encontrada.", url.String()),
		})
	})
}
