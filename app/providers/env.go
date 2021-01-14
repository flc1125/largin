package providers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Env(e *gin.Engine) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
