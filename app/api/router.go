package api

import (
	"github.com/gin-gonic/gin"
	"orse/internal/auth"
	"orse/internal/menu"
	"orse/internal/middleware"
	"orse/internal/user"
)

func Api(r *gin.Engine)  {

	v1 := r.Group("/v1")
	{
		v1.POST("/add-casbin",auth.AddCasbin)

		v1.POST("/editMenu",menu.EditMenu)
		v1.POST("/get-menu",menu.GetMenu)
		v1.POST("/user/register",user.Register)
		v1.POST("/user/login",user.GetToken)
		v1.POST("/user/get-token-user",user.GetTokenUser)


		v1.POST("/user/r-token",user.RefreshToken)


		users := v1.Group("/user")


		users.Use(middleware.JWThMiddleware())
		////用户中心
		users.GET("/index", user.Info)
		users.GET("/address/list", user.Info)
		////消息通知
		//users.GET("/msgCenter",login)

	}
}
