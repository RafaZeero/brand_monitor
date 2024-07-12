package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
