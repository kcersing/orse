package middleware

import (
	"github.com/gin-gonic/gin"
	"orse/internal/jwt"
	"orse/internal/util"
)

// JWThMiddleware 中间件
func JWThMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			// 处理 没有token的时候
			util.Error(c, 1,"缺少token信息")
			c.Abort() // 不会继续停止
			return
		}
		// 解析
		j := jwt.NewJWT()
		claims, err := j.VerifyToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				util.Error(c, 1,"授权已过期")
				c.Abort()
				return
			}
			// 处理 解析失败
			util.Error(c, 1, err.Error())
			c.Abort()
			return
		}
		c.Set("claims", claims)
		AuthCheckRole()
		c.Next()
	}
}
