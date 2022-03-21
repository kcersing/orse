package middleware

import (
	"github.com/gin-gonic/gin"
	"orse/internal/auth"
	"orse/internal/jwt"
	"orse/internal/util"
)

// AuthCheckRole 权限检查
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(*jwt.Claims)
		role := claims.Role
		//e := auth.CasbinEnforcer()
		// 检查权限
		ok, err := auth.CasbinEnforcer.Enforce(role, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			util.Error(c, 1, err.Error())
			c.Abort()
			return
		}

		if ok == true {
			c.Next()
		} else {
			util.Error(c, 1, "很抱歉您没有此权限")
			c.Abort()
			return
		}
	}
}
