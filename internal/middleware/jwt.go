package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orse/internal/jwt"
)

// JWThMiddleware 中间件
func JWThMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			// 处理 没有token的时候
			c.JSON(http.StatusOK, gin.H{
				"message": "缺少token信息",
				"code":    401,
			})
			c.Abort() // 不会继续停止
			return
		}
		// 解析
		j := jwt.NewJWT()
		claims, err := j.VerifyToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": 401,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			// 处理 解析失败
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
				"code":    401,
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		AuthCheckRole()
		c.Next()
	}
}
