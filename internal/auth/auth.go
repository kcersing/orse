package auth

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2/bson"
	"orse/internal/util"
)
func AddRoleForUser(c *gin.Context) {
	ca := CasbinModel{
		ID:     bson.NewObjectId(),
		Uid:    "1",
		Role:   "role",
		Path:   "",
		Method: "",
	}
	_, err := ca.AddRoleForUser()
	if err != nil {
		util.Error(c, 1, err.Error())
		return
	}
	util.Success(c, nil)
	return
}

func grantUserRole() {

}

func VerifyCasbin(c *gin.Context) {

}
