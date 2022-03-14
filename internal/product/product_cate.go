package product

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"orse/internal/errors"
)

type cateParams struct {
	Id   int    `form:"Id" json:"Id" `
	Name string `form:"name" json:"name" binding:"required"`
	Sort string `form:"sort" json:"sort" binding:"required" `
	Pid  int    `form:"pid" json:"pid" binding:"required" `
}

type CateResult struct {
	Id   int    `form:"Id" json:"Id" `
	Name string `form:"name" json:"name" binding:"required"`
	Sort string `form:"sort" json:"sort" binding:"required" `
	Pid  int    `form:"pid" json:"pid" binding:"required" `
}
func EditCate(c *gin.Context) {
	var ct cateParams
	if err := c.ShouldBind(&ct); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusOK, gin.H{
				"message": util.GetError(fe),
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
	cateParams.EditCate(ct)
	return
}

func (c *cateParams)EditCate() {

}



