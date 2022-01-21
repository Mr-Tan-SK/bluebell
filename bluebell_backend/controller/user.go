package controller

import (
	"bluebell_backend/logic"
	"bluebell_backend/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 注册业务
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	u := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(u); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ServerResponse(c, CodeInvalidParam, nil)
		} else {
			msg := removeTopStruct(errs.Translate(trans))
			ServerResponseWithMsg(c, CodeInvalidParam, msg, nil)
		}
		return
	}
	// 2. 业务处理
	if err := logic.Signup(u); err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		ServerResponse(c, CodeUserExist, nil)
		return
	}
	// 3. 返回响应
	ServerResponse(c, CodeSuccess, nil)
}

func LoginHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	u := new(models.ParamLogin)
	if err := c.ShouldBindJSON(u); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ServerResponse(c, CodeInvalidParam, nil)
		} else {
			msg := removeTopStruct(errs.Translate(trans))
			ServerResponseWithMsg(c, CodeInvalidParam, msg, nil)
		}
		return
	}
	// 2. 业务处理
	token, err := logic.Login(u)
	if err != nil {
		zap.L().Error("登录失败", zap.Error(err))
		ServerResponseWithMsg(c, CodeInvalidPassword, err.Error(), nil)
		return
	}
	// 3. 返回响应
	ServerResponse(c, CodeSuccess, token)
}

func RefreshTokenHandler(c *gin.Context) {

}
