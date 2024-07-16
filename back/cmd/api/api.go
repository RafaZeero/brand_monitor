package api

import (
	"fmt"
	"net/http"
	"os"

	googlesearch "github.com/RafaZeero/brand_monitor/services/google-search"
	"github.com/RafaZeero/brand_monitor/services/healthz"
	"github.com/RafaZeero/brand_monitor/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type ApiServer struct {
	addr string
	db   *mongo.Client
}

func NewApiServer(addr string, db *mongo.Client) *ApiServer {
	return &ApiServer{addr: addr, db: db}
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
	s.RegisterRoutes(r)

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

func (s *ApiServer) RegisterRoutes(r *gin.Engine) {
	healthzStore := healthz.NewStore(s.db.Database("test").Collection("test"))
	healthzHandler := healthz.NewHandler(healthzStore)

	searchStore := googlesearch.NewStore(s.db.Database("searches").Collection("competitors"))
	searchHandler := googlesearch.NewHandler(searchStore)

	v1 := r.Group("/v1")
	{
		healthzHandler.RegisterRoutes(v1)
		searchHandler.RegisterRoutes(v1)
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
