package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func RespondWithJSON(ctx *gin.Context, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal JSON response: %v", payload)
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(code)
	ctx.Writer.Write(data)
}

func RespondWithError(ctx *gin.Context, code int, msg string) {
	if code > 499 {
		log.Printf("respond with 5XX error: %v", msg)
	}

	type errResponse struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}

	RespondWithJSON(ctx, code, errResponse{Error: msg, Success: false})
}

func EnvOrFatal(key string) string {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", err))
	}

	value := os.Getenv(key)
	if value == "" {
		panic("missing required environment variable " + key)
	}

	return value
}
