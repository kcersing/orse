package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"orse/common/models"
	"orse/internal/database"
	"orse/internal/jwt"
	"orse/internal/util"
)

type LoginRequest  struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required" `
	Pass   string `form:"password" json:"password" binding:"required" `
}
type VerifyTokenUser struct{
	Token  string `form:"token" json:"token" binding:"required" `
}

func GetToken(c *gin.Context) {
	var l LoginRequest
	if err := c.ShouldBind(&l); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			util.Error(c,1,util.GetError(fe))
		}
		return
	}

	client := database.Open()
	//defer client.Close()
	u:=models.User{}
	result := client.Where("mobile = ?", l.Mobile).First(&u)
	if result.Error != nil {
		util.Error(c,1,"手机号不存在")
		return
	}
	if !VerifyPass(u.Pass, []byte(l.Pass)) {
		util.Error(c,1,"用户名或密码错误")
		return
	}
	//生成token
	j := jwt.NewJWT()
	tokenString, err := j.GenToken(	jwt.NewClaims(int(u.Id), int(u.Role),u.Username))
	if err != nil {
		util.Error(c,1, err.Error())
		return
	}
	util.Success(c,gin.H{"token": tokenString})
	return
}

func GetTokenUser(c *gin.Context){
	var l VerifyTokenUser
	if err := c.ShouldBind(&l); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			util.Error(c,1, util.GetError(fe))
		}
		return
	}
	//生成token
	j := jwt.NewJWT()
	claims, err := j.VerifyToken(l.Token)
	if err != nil {
		util.Error(c,1, err.Error())
		return
	}
	util.Success(c,claims)
	return
}


type refreshToken struct {
	Token string `form:"token" json:"token" binding:"required" `
}

func RefreshToken(c *gin.Context) {
	var k refreshToken
	if err := c.ShouldBind(&k); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			util.Error(c,1, util.GetError(fe))
			return
		}
	}

	if k.Token == "" {
		k.Token = c.GetHeader("token")
	}

	if k.Token == "" {
		util.Error(c,1,"token不存在")
		return
	}

	j := jwt.NewJWT()
	tokenString, err := j.RefreshToken(k.Token)
	if err != nil {
		util.Error(c,1,err.Error())
		return
	}
	util.Success(c,gin.H{"token": tokenString})
	return
}
