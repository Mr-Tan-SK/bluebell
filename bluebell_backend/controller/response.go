package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

// ServerResponse 服务器返回的响应信息
func ServerResponse(c *gin.Context, code ResCode, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  code.Msg(),
		Data: data,
	})
}

// ServerResponseWithMsg 服务器返回自定义Msg
func ServerResponseWithMsg(c *gin.Context, code ResCode, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
