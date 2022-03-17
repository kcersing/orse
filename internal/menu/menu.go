package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"orse/common/models"
	"orse/internal/database"
	"orse/internal/util"
)

type Menu struct {
	Id       int    `form:"Id" json:"Id" `
	Title    string `form:"title" json:"title" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required" `
	Url      string `form:"url" json:"url" binding:"required" `
	ParentId int    `form:"parent_id" json:"parent_id" binding:"required" `
}

func (menu *Menu) EditMenu(c *gin.Context) {
	client, _ := database.Open()
	var menus models.Menu
	if menu.ParentId != 0 {
		client.First(&menus)
		menus.Title = menu.Title
		menus.Name = menu.Name
		menus.Url = menu.Url
		menus.ParentId = int64(menu.ParentId)
		result := client.Save(&menus)
		if result.Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": result.Error,
				"code":    1,
			})
			return
		}
	} else {
		createMenu := models.Menu{Name: menu.Name, Title: menu.Title, Url: menu.Url, ParentId: int64(menu.ParentId)}
		result := client.Create(&createMenu)
		if result.Error != nil {
			util.Error(c,1,result.Error.Error())
			return
		}
	}
	util.Success(c , nil)
	return
}

func EditMenu(c *gin.Context) {
	var menu Menu
	if err := c.ShouldBind(&menu); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			util.Error(c,1,util.GetError(fe))
		} else {
			util.Error(c,1,err.Error())
		}
		return
	}
	menu.EditMenu(c)
	return
}

type AllMenu struct {
	Id        int
	Title     string
	Name      string
	Url       string
	ParentId  int
	Component string
	Icon      string
}

type AllMenuNode struct {
	Id        int
	Title     string
	Name      string
	Url       string
	ParentId  int
	Component string
	Icon      string
	Children  []AllMenuNode
}

func getChildNode(m []AllMenu, pid int) []AllMenuNode {
	tree := make([]AllMenuNode, 0)
	for _, item := range m {
		if pid == item.ParentId {
			node := AllMenuNode{
				Id:        item.Id,
				Title:     item.Title,
				Name:      item.Name,
				Url:       item.Url,
				ParentId:  item.ParentId,
				Component: item.Component,
				Icon:      item.Icon,
				Children:  nil,
			}
			node.Children = getChildNode(m, node.Id)

			tree = append(tree, node)
		}
	}
	return tree
}

func GetMenu(c *gin.Context) {
	client, _ := database.Open()
	var lists []AllMenu
	result := client.Model(&models.Menu{}).Find(&lists)

	tree := getChildNode(lists, 0)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": result.Error,
			"code":    1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"data":    tree,
		"code":    0,
	})
	return
}
