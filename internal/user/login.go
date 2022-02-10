package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"orse/ent/user"
	"orse/internal/database"
	"orse/internal/errors"
	"orse/internal/jwt"
)

type LoginRequest  struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required" `
	Pass   string `form:"pass" json:"pass" binding:"required" `
}
type  LoginResponse struct{
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}
func GetTokenHandler(c *gin.Context) {
	var l LoginRequest
	if err := c.ShouldBind(&l); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errors.GetError(fe),
				"code":    1,
			})
		}
		return
	}

	client, _ := database.Open()
	defer client.Close()

	u, err := client.User.
		Query().
		Where(user.Mobile(l.Mobile)).
		Only(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "手机号不存在",
			"code":    1,
		})
		return
	}
	if !VerifyPass(u.Pass, []byte(l.Pass)) {
		c.JSON(http.StatusOK, gin.H{
			"code": 202,
			"msg":  "用户名或密码错误",
		})
		return
	}

	//生成token
	tokenString, err := jwt.GenToken(u.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 101,
			"err":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": tokenString},
	})
	return

}

type token struct {
	Token string `form:"token" json:"token" `
	Id    int    `form:"id" json:"id" binding:"required" `
}

func VerifyToken(c *gin.Context) {
	var k token
	if err := c.ShouldBind(&k); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errors.GetError(fe),
				"code":    1,
			})
			return
		}
	}

	l, err := jwt.VerifyToken(k.Token)
	log.Println(l, err)

}

func RefreshToken(c *gin.Context) {
	var k token
	if err := c.ShouldBind(&k); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errors.GetError(fe),
				"code":    1,
			})
			return
		}
	}

	if k.Token == "" {
		k.Token = c.GetHeader("token")
	}

	if k.Token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token不存在",
			"code":    1,
		})
		return
	}

	tokenString, err := jwt.RefreshToken(k.Id, k.Token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    1,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"data": gin.H{ "token": tokenString },
		"code": 0,
	})
	return
}
