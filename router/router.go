package router

import (
	"github.com/gin-gonic/gin"
)

var R *gin.Engine
func init() {
	R = gin.Default()
	R.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"code": 400, "message": "Bad Request"})
	})

}




