package api

import (
	"github.com/gin-gonic/gin"
	"orse/internal/auth"
	"orse/internal/jwt"
	"orse/internal/menu"
	"orse/internal/user"
)

func Api(r *gin.Engine)  {

	v1 := r.Group("/v1")
	{

		v1.GET("/login",login)
		v1.POST("/roles",auth.Roles)
		v1.POST("/roleAuths",auth.RoleAuths)
		v1.POST("/addMenu",menu.AddMenu)
		v1.POST("/user/register",user.Register)
		v1.POST("/user/login",user.GetToken)
		v1.POST("/user/r-token",user.RefreshToken)


		users := v1.Group("/user")
		users.Use(jwt.JWThMiddleware())
		//用户中心
		users.GET("/index", user.Info)
		//消息通知
		users.GET("/msgCenter",login)
	//	info 账号信息
	//address/list 地址管理

	}



}
func login(c *gin.Context)  {
	c.JSON(200,gin.H{
		"data":"hello",
		"message":"hh",
	})
}

