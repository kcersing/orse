package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"orse/common/models"
	"orse/internal/database"
	"orse/internal/errors"
)

type Menu struct {
	Id       int    `form:"Id" json:"Id" `
	Title    string `form:"title" json:"title" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required" `
	Url      string `form:"url" json:"url" binding:"required" `
	ParentId int    `form:"parent_id" json:"parent_id" binding:"required" `
}
var menus models.Menu
func (menu *Menu) EditMenu(c *gin.Context) {
	client, _ := database.Open()

	if menu.ParentId != 0 {
		client.First(&menus)
		menus.Title = menu.Title
		menus.Name = menu.Name
		menus.Url = menu.Url
		menus.ParentId = int64(menu.ParentId)
		result :=client.Save(&menus)
		if result.Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": result.Error,
				"code":    1,
			})
			return
		}
	} else {
		createMenu := models.Menu{Name:  menu.Name, Title: menu.Title, Url: menu.Url,ParentId:int64(menu.ParentId)}
		result := client.Create(&createMenu)
		if result.Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": result.Error,
				"code":    1,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"code":    0,
	})
	return
}

func EditMenu(c *gin.Context) {
	var menu Menu
	if err := c.ShouldBind(&menu); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusOK, gin.H{
				"message": errors.GetError(fe),
				"code":    1,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": err,
				"code":    1,
			})
		}
		return
	}
	menu.EditMenu(c)
	return
}

func GetMenu(c *gin.Context) {
	client, _ := database.Open()

	result := client.Select("id", "title","name","url","parent_id","component","icon").Find(&menus)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": result.Error,
			"code":    1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"data":    menus,
		"code":    0,
	})
	return
}
