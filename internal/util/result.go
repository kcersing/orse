package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//返回的结果：
type Result struct {
	Time  time.Time   `json:"time"`
	Code  int         `json:"code"`
	Message   string      `json:"msg"`
	Data  interface{} `json:"data"`
}


//成功
func Success(c *gin.Context, data interface{}) {
	if (data == nil) {
		data = gin.H{}
	}
	res := Result{
		Time:    time.Now(),
		Code:    200,
		Message: "成功",
		Data:    data,
	}
	c.JSON(http.StatusOK,res)
}

//出错
func Error(c *gin.Context, code int,message string) {
	res := Result{
		Time:    time.Now(),
		Code:    code,
		Message: message,
		Data:    gin.H{},
	}
	c.JSON(http.StatusOK,res)
}
