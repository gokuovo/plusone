package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 封装了API的统一响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Success 发送一个成功的响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  "操作成功",
		Data: data,
	})
}

// Error 发送一个失败的响应
func Error(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, Response{
		Code: http.StatusInternalServerError,
		Msg:  err.Error(),
		Data: nil,
	})
}
