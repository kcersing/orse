package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"saas/internal/database"
	"saas/internal/errors"
)

type Menu struct {
	Id int  `form:"Id" json:"Id" `
	Title string `form:"title" json:"title" binding:"required"`
	Name  string `form:"name" json:"name" binding:"required" `
	Url   string `form:"url" json:"url" binding:"required" `
	Pid   int  `form:"pid" json:"pid" binding:"required" `
}


func AddMenu(c *gin.Context) {
	client,_ := database.Open()
	defer client.Close()
	var menu Menu
	if err := c.ShouldBind(&menu); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errors.GetError(fe),
				"code":    1,
			})
		}else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
				"code":    1,
			})
		}
		return
	}

	model, err := client.Menu.
		Create().
		SetTitle(menu.Title).
		SetName(menu.Name).
		SetURL(menu.Url).
		SetPid(menu.Pid).
		SetLevel(1).
		SetSort(1).
		SetTree("1").
		SetIcon("111").
		SetStatus("1").
		SetDesc("1").
		Save(c)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
			"code":    1,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"data":    model,
		"code":    0,
	})
	return

}

func editMenu(c *gin.Context) {
	client,_ := database.Open()
	defer client.Close()
	var menu Menu

	if err := c.ShouldBind(&menu); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errors.GetError(fe),
				"code":    1,
			})
		}else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
				"code":    1,
			})
		}
		return
	}

	model, err := client.Menu.UpdateOneID(menu.Id).
		SetTitle(menu.Title).
		SetName(menu.Name).
		SetURL(menu.Url).
		SetPid(menu.Pid).
		SetLevel(1).
		SetSort(1).
		SetTree("1").
		SetIcon("111").
		SetStatus("1").
		SetDesc("1").
		Save(c)

	//tr_0 tr_5

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
			"code":    1,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"data":    model,
		"code":    0,
	})
	return
}

func  GetMene()  {
	client,_ := database.Open()
	defer client.Close()

}