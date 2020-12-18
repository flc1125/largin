package users

import (
	"github.com/flc1125/go-gin/app/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/**
 * 用户 API
 */
func Info(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("databases/gin.db"), &gorm.Config{})

	if err != nil {
		panic("database connect error")
	}

	db.Create(&models.User{User: "tests"})

	c.JSON(200, gin.H{
		"user": "FLC223",
	})
}
