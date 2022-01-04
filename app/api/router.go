package api

import (
	"github.com/gin-gonic/gin"
	"orse/internal/auth"
	"orse/internal/menu"
)

func Api(r *gin.Engine)  {
	v1 :=r.Group("/v1")
	{
		v1.GET("/login",login)
		v1.POST("/roles",auth.Roles)
		v1.POST("/role-auths",auth.RoleAuths)
		v1.POST("/add-menu",menu.AddMenu)

	}

}
func login(c *gin.Context)  {
	c.JSON(200,gin.H{
		"data":"hello",
		"message":"hh",
	})
}

