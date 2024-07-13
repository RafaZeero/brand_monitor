package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

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
	// get root dir
	rootDir := getRootDir()

	envPath := filepath.Join(rootDir, "/back/.env")

	// load .env file
	if err := godotenv.Load(envPath); err != nil {
		panic(fmt.Sprintf("Error loading .env file - path: %s", envPath))
	}

	value := os.Getenv(key)
	if value == "" {
		panic("missing required environment variable " + key)
	}

	return value
}

func getRootDir() string {
	// get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// search for the root directory by looking for the .git folder
	for {
		if _, err := os.Stat(filepath.Join(cwd, ".git")); err == nil {
			// found the .git folder, return this directory as the root
			return cwd
		}

		// move to the parent directory
		parent := filepath.Dir(cwd)
		if parent == cwd {
			// reached the root without finding .git folder, return current directory
			return cwd
		}
		cwd = parent
	}
}
