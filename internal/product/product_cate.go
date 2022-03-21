package product

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"orse/common/models"
	"orse/internal/database"
	"orse/internal/util"
)

type cateParams struct {
	Id   int64    `form:"Id" json:"Id" `
	Name string `form:"name" json:"name" binding:"required"`
	Sort int64 `form:"sort" json:"sort" binding:"required" `
	Pid  int64    `form:"pid" json:"pid" binding:"required" `
}

type CateResult struct {
	Id   int64    `form:"Id" json:"Id" `
	Name string `form:"name" json:"name" binding:"required"`
	Sort int64 `form:"sort" json:"sort" binding:"required" `
	Pid  int64    `form:"pid" json:"pid" binding:"required" `
}
func EditCate(c *gin.Context) {
	var ct cateParams
	if err := c.ShouldBind(&ct); err != nil {
		if fe, ok := err.(validator.ValidationErrors); ok {
			util.Error(c,1, util.GetError(fe))
		} else {
			util.Error(c,1, err.Error())
		}
		return
	}
	ct.EditCate(c)
	return
}

func (cate *cateParams)EditCate(c *gin.Context) {

	client := database.Open()
	var cates models.ProductCate
	if cate.Id != 0 {
		cates.Id = cate.Id
		client.First(&cates)
		cates.ParentId = cate.Pid
		cates.Name = cate.Name
		cates.Sort = cate.Sort
		result := client.Save(&cates)
		if result.Error != nil {
			util.Error(c,1,result.Error.Error())
			return
		}
	} else {
		createProductCate := models.ProductCate{
			ParentId:     cate.Pid,
			Status:       0,
			Name:        cate.Name,
			Sort:         cate.Sort,
		}
		result := client.Create(&createProductCate)
		if result.Error != nil {
			util.Error(c, 1, result.Error.Error())
			return
		}
	}
	util.Success(c , nil)
	return
}

