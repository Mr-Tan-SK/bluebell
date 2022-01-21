package logic

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
)

func GetCommunityList() (data interface{}, err error) {
	// 1. 查询数据库, 找到所有的community并返回
	data, err = mysql.GetCommunityList()
	return
}

func GetCommunityDetail(id int64) (data models.CommunityDetail, err error) {
	data, err = mysql.GetCommunityById(id)
	return
}
