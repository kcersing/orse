package auth

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2/bson"
)



//func AddCasbin(c *gin.Context) {
//	role, _ := c.GetPostForm("role")     // 想要访问资源的用户。
//	path, _ := c.GetPostForm("path")     // 将被访问的资源。
//	method, _ := c.GetPostForm("method") // 用户对资源执行的操作。
//	ptype := "p"
//	ca := CasbinModel{
//		ID:     bson.NewObjectId(),
//		Ptype:  ptype,
//		Role:   role,
//		Path:   path,
//		Method: method,
//	}
//	_, err := ca.addPolicy()
//	if err != nil {
//
//		util.Error(c, 1, "保存失败")
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code":    200,
//		"message": "保存成功",
//	})
//	util.Success(c, nil)
//}

func AddRole(c *gin.Context) {
	//role, _ := c.GetPostForm("role")     // 想要访问资源的用户。

	ca := CasbinModel{
		ID:     bson.NewObjectId(),
		Uid: 1,
		Role:   "role",
	}
 	ca.AddRoleForUser()

}

func grantUserRole() {

}

func VerifyCasbin(c *gin.Context) {

}
