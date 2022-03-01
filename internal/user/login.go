package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"orse/common/models"

	"orse/internal/database"
	"orse/internal/errors"
	"orse/internal/jwt"
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
			c.JSON(http.StatusOK, gin.H{
				"code":    201,
				"message": errors.GetError(fe),
			})
		}
		return
	}

	client, _ := database.Open()
	//defer client.Close()
	u:=models.User{}
	result := client.Where("mobile = ?", l.Mobile).First(&u)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    201,
			"message": "手机号不存在",
		})
		return
	}
	if !VerifyPass(u.Pass, []byte(l.Pass)) {
		c.JSON(http.StatusOK, gin.H{
			"code": 202,
			"message":  "用户名或密码错误",
		})
		return
	}
	//生成token
	j := jwt.NewJWT()
	tokenString, err := j.GenToken(	jwt.NewClaims(int(u.Id), int(u.Role),u.Username))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 201,
			"message":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message":    "登录成功",
		"data": gin.H{"token": tokenString},
	})
	return
}

func GetTokenUser(c *gin.Context){
	var l VerifyTokenUser
	if err := c.ShouldBind(&l); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusOK, gin.H{
				"code":    201,
				"message": errors.GetError(fe),
			})
		}
		return
	}
	//生成token
	j := jwt.NewJWT()
	claims, err := j.VerifyToken(l.Token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 201,
			"message":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message":    "获取成功",
		"data": claims,
	})
	return
}


type refreshToken struct {
	Token string `form:"token" json:"token" binding:"required" `
}

func RefreshToken(c *gin.Context) {
	var k refreshToken
	if err := c.ShouldBind(&k); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusOK, gin.H{
				"code":    201,
				"message": errors.GetError(fe),
			})
			return
		}
	}

	if k.Token == "" {
		k.Token = c.GetHeader("token")
	}

	if k.Token == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    201,
			"message": "token不存在",
		})
		return
	}

	j := jwt.NewJWT()
	tokenString, err := j.RefreshToken(k.Token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"code":    1,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{ "token": tokenString },
		"code": 0,
	})
	return
}
