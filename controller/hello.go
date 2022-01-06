package hello

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	msg := "Hello World!"
	fmt.Println(c.MustGet("example"))	// example参数来自middlerware
	c.String(http.StatusOK, msg)
}