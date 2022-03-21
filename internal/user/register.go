package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"orse/common/models"
	"orse/internal/database"
	"orse/internal/util"
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
			util.Error(c,1,util.GetError(fe))
			return
		}
	}

	client := database.Open()

	result := client.Where("mobile = ?", r.Mobile).First(&u)

	if result.Error == nil {
		util.Error(c,1,"手机号已存在")
		return
	}
	user := models.User{
		Username: r.Username,
		Pass:     SetPass([]byte(r.Pass)),
		Mobile:   r.Mobile,
	}

	err := client.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&user).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Create(&models.UserDetail{
			UserId: user.Id,
			Name:   r.Username,
		}).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})

	if err != nil {
		util.Error(c,1,err.Error())
		return
	}
	util.Success(c,nil)
	return
}
