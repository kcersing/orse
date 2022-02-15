package auth

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"orse/internal/database"
	"orse/internal/jwt"
)

type CasbinModel struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Ptype  string        `json:"ptype" bson:"ptype"`
	Role   string        `json:"role" bson:"v0"`
	Path   string        `json:"path" bson:"v1"`
	Method string        `json:"method" bson:"v2"`
}

func CasbinEnforcer() *casbin.Enforcer {
	a, _ := database.Open()
	e, err := casbin.NewEnforcer("config/casbin/casbin_model.ini", a)
	if err != nil {
		panic(err)
	}
	e.LoadPolicy()
	return e
}
func (c *CasbinModel) AddCasbin(cm CasbinModel) (bool, error) {
	n := CasbinEnforcer()
	return n.AddPolicy(cm.Role, cm.Path, cm.Method)
}

func AddCasbin(c *gin.Context) {

	role, _ := c.GetPostForm("role")     // 想要访问资源的用户。
	path, _ := c.GetPostForm("path")     // 将被访问的资源。
	method, _ := c.GetPostForm("method") // 用户对资源执行的操作。

	ptype := "p"
	ca := CasbinModel{
		ID:     bson.NewObjectId(),
		Ptype:  ptype,
		Role:   role,
		Path:   path,
		Method: method,
	}
	_, err := ca.AddCasbin(ca)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    201,
			"message": "保存失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "保存成功",
	})
}

// AuthCheckRole 权限检查
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(*jwt.Claims)
		role := claims.Role
		e := CasbinEnforcer()
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
