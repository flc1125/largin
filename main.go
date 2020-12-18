package main

import (
	"github.com/flc1125/go-gin/app/providers"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	providers.Application(e)

	e.Run()
}
