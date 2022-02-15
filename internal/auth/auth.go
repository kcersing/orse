package auth

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"orse/internal/database"
)

func main() {
	a,_ := database.Open()
	e, err := casbin.NewEnforcer("configs/casbin_model.ini", a)
	if err != nil{
		panic(err)
	}

	dded,_:=e.AddPolicy("alice","dara","read")
	log.Println(dded)
	Policy:=e.GetPolicy()
 	log.Println(Policy)

	GetPolicy(e, [][]string{{"eve", "data3", "read"},{"alice","dara","read"}})

}
func Roles(c *gin.Context)  {

	sub,_:=c.GetPostForm("sub")// 想要访问资源的用户。
	obj,_:=c.GetPostForm("obj")// 将被访问的资源。
	act,_:=c.GetPostForm("act") // 用户对资源执行的操作。
	e := newEnforcer()
	ok,err := e.AddPolicy(sub,obj,act)

	if err != nil{
		c.JSON(200,gin.H{
			"message":err,
			"code":1,
		})
	}
	if ok{
		c.JSON(200,gin.H{
			"message":"成功",
			"code":0,
		})
	}
}

func DelRoles(c *gin.Context){

}

func RoleAuths(c *gin.Context)   {
	sub,_:=c.GetPostForm("sub")// 想要访问资源的用户。
	obj,_:=c.GetPostForm("obj")// 将被访问的资源。
	act,_:=c.GetPostForm("act") // 用户对资源执行的操作。

	e := newEnforcer()
	ok,err := e.AddPolicy(sub,obj,act)

	if err != nil{
		c.JSON(200,gin.H{
			"message":err,
			"code":1,
		})
	}
	if ok{
		c.JSON(200,gin.H{
			"message":"成功",
			"code":0,
		})
	}
}

func newEnforcer() *casbin.Enforcer{
	a,_ := database.Open()
	e, err := casbin.NewEnforcer("configs/                                    .ini", a)
	if err != nil{
		panic(err)
	}
	return e
}

func GetPolicy(e *casbin.Enforcer, res [][]string) {
	myRes := e.GetPolicy()
	log.Println("Policy: ", myRes)

	if !util.Array2DEquals(res, myRes) {
		log.Fatalf("Policy: %v, supposed to be %v", myRes, res)
	}}

