package product

import (
	"github.com/gin-gonic/gin"
)

type Cate struct {
	Id   int    `form:"Id" json:"Id" `
	Name string `form:"name" json:"name" binding:"required"`
	Sort string `form:"sort" json:"sort" binding:"required" `
	Pid  int    `form:"pid" json:"pid" binding:"required" `
}

func CreateCate(ctx *gin.Context) {

}
