package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"orse/ent/user"
	"orse/internal/database"

	"net/http"
	"orse/internal/errors"
)

type profile struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Username string `form:"username" json:"username" binding:"required" `
	Pass     string `form:"pass" json:"pass" binding:"required" `
}

func Register(c *gin.Context) {
	var p profile
	if err := c.ShouldBind(&p); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errors.GetError(fe),
				"code":    1,
			})
			return
		}
	}

	client, _ := database.Open()
	defer client.Close()

	_, err := client.User.
		Query().
		Where(user.Mobile(p.Mobile)).
		Only(c)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "手机号已存在",
			"code":    1,
		})
		return
	}

	tx, _ := client.Tx(c)

	xu, err := tx.User.
		Create().
		SetUsername(p.Username).
		SetPass(SetPass([]byte(p.Pass))).
		SetMobile(p.Mobile).
		Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"code":    0,
		})
		return
	}

	_, err = tx.UserDetail.
		Create().
		SetName(xu.Username).
		SetUser(xu).
		Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"code":    0,
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "注册成功",
		"code":    0,
	})
	return
}
