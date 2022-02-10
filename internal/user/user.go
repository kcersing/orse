package user

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"orse/ent/user"
	"orse/internal/database"
)

func Info(c *gin.Context) {
	userId, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "userID值不存在",
			"code":    401,
		})
		return
	}
	client, _ := database.Open()
	defer client.Close()
	u, err := client.User.
		Query().
		Where(user.ID(userId.(int))).
		Only(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "会员不存在",
			"code":    201,
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 0,
		"data": u,
	})

	log.Println(userId)
	return
}

// SetPass 设置密码
func SetPass(pwd []byte) string {
	pass, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(pass)
}

// VerifyPass 验证密码
func VerifyPass(pass string, pad []byte) bool {
	byteHash := []byte(pass)
	err := bcrypt.CompareHashAndPassword(byteHash, pad)
	if err != nil {
		return false
	}
	return true
}
