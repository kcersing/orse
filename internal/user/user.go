package user

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"orse/internal/database"
	"orse/internal/util"
)

func Info(c *gin.Context) {
	userId, ok := c.Get("userID")
	if !ok {
		util.Error(c, 1,"userID值不存在")
		return
	}
	client := database.Open()

	result := client.Where("id = ?", userId.(int)).First(&u)

	if result.Error != nil {
		util.Error(c, 1,"会员不存在")
		return
	}
	util.Success(c, u)
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
