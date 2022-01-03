package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {

}

func Router() *gin.Engine  {
	r:=gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"hello world")
	})

	return r

}





