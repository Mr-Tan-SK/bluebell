package logic

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
	"bluebell_backend/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 生成post id
	p.PostID, err = snowflake.GenID()
	// 2. 保存到数据库中
	err = mysql.CreatePost(p)
	// 3. 返回
	return
}

func GetPostById(id int64) (data *models.Post, err error) {
	data, err = mysql.GetPostById(id)
	if err != nil {

	}
	return
}
