package payments

import (
	"os/exec"

	"github.com/gin-gonic/gin"
)

func Alipay(c *gin.Context) {
	cmd := exec.Command("ls", "/")
	output, _ := cmd.CombinedOutput()

	c.JSON(200, gin.H{
		"payments": "alipay",
		"ret":      output,
	})
}
