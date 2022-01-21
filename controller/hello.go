package hello

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"test/internal/config"
)

func Hello(c *gin.Context) {
	// msg := "Hello World!"
	fmt.Println(c.MustGet("example"))	// example参数来自middlerware
	c.String(http.StatusOK, config.Conf.Mysql.Addr)
}