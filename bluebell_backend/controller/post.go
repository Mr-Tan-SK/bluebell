package controller

import (
	"bluebell_backend/logic"
	"bluebell_backend/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数的校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CreatePostHandler failed", zap.Error(err))
		ServerResponse(c, CodeInvalidParam, nil)
		return
	}
	userID, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID failed", zap.Error(err))
		return
	}
	p.AuthorID = userID
	// 2. 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ServerResponse(c, CodeInvalidParam, nil)
		return
	}
	// 3. 返回响应
	ServerResponse(c, CodeSuccess, p)
}

func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数,(从URL中获取帖子的id)
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("invalid param", zap.Error(err))
		ServerResponse(c, CodeInvalidParam, nil)
		return
	}

	// 2. 根据id取出帖子的数据(查数据库)
	data, err := logic.GetPostById(id)
	if err != nil {
		zap.L().Error("logic.GetPostById(id) failed", zap.Error(err))
		ServerResponse(c, CodeServerBusy, nil)
		return
	}
	// 3. 返回响应
	ServerResponse(c, CodeSuccess, data)
}
