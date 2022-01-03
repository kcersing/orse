package auth

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	"github.com/casbin/ent-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

func main() {
	a := initAdapter()
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


func initAdapter() *entadapter.Adapter {
	cfg, err := ini.Load("configs/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	dataType := cfg.Section("database").Key("TYPE").String()
	dataUser := cfg.Section("database").Key("USER").String()
	dataPassword := cfg.Section("database").Key("PASSWORD").String()
	dataHost := cfg.Section("database").Key("HOST").String()
	dataName := cfg.Section("database").Key("NAME").String()

	a, err := entadapter.NewAdapter(dataType, dataUser+":"+dataPassword+"@tcp("+dataHost+")/"+dataName)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	return a
}

func newEnforcer() *casbin.Enforcer{
	a := initAdapter()
	e, err := casbin.NewEnforcer("configs/casbin_model.ini", a)
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
	}
}
