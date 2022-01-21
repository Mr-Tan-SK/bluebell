package controller

import (
	"bluebell_backend/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CommunityHandler(c *gin.Context) {
	// 1. 查询到所有的社区(id, name)以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ServerResponse(c, CodeServerBusy, nil)
		return
	}
	ServerResponse(c, CodeSuccess, data)
}

func CommunityDetailHandler(c *gin.Context) {
	// 1. 获取社区id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("logic getCommunityId failed", zap.Error(err))
		ServerResponse(c, CodeInvalidParam, nil)
		return
	}
	// 1.根据id查询社区详细信息
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic GetCommunityDetail failed", zap.Error(err))
		ServerResponse(c, CodeInvalidParam, nil)
		return
	}
	ServerResponse(c, CodeSuccess, data)
}
