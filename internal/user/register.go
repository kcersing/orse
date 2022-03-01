package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"orse/common/models"
	"orse/internal/database"
	"orse/internal/errors"
)

type RegisterRequest struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Username string `form:"username" json:"username" binding:"required" `
	Pass     string `form:"pass" json:"pass" binding:"required" `
}
var u = models.User{}
func Register(c *gin.Context) {
	var r RegisterRequest
	if err := c.ShouldBind(&r); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errors.GetError(fe),
				"code":    1,
			})
			return
		}
	}

	client, _ := database.Open()

	result := client.Where("mobile = ?", r.Mobile).First(&u)

	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "手机号已存在",
			"code":    1,
		})
		return
	}

	user := models.User{
		Username : r.Username,
		Pass : SetPass([]byte(r.Pass)),
		Mobile : r.Mobile,

	}
	result = client.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": result.Error,
			"code":    1,
		})
		return
	}
	result = client.Create(&models.UserDetail{
		UserId:user.Id,
		Name: r.Username,


	})
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": result.Error,
			"code":    1,
		})
		return
	}




	c.JSON(http.StatusBadRequest, gin.H{
		"message": "注册成功",
		"code":    0,
	})
	return
}
