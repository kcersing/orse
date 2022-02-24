package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orse/internal/auth"
	"orse/internal/jwt"
)

// AuthCheckRole 权限检查
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(*jwt.Claims)
		role := claims.Role
		e := auth.CasbinEnforcer()
		// 检查权限
		ok, err := e.Enforce(role, c.Request.URL.Path, c.Request.Method)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  201,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		if ok == true {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  201,
				"message": "很抱歉您没有此权限",
			})
			c.Abort()
			return
		}
	}
}
