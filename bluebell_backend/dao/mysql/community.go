package mysql

import (
	"bluebell_backend/models"
	"database/sql"
	"errors"
)

func GetCommunityList() (data []models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&data, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			data, err = []models.Community{}, nil
		}
		return nil, errors.New("community Select failed, err: " + err.Error())
	}
	return
}

func GetCommunityById(id int64) (data models.CommunityDetail, err error) {
	sqlStr := "select community_id, community_name, introduction, create_time from community where community_id=?"
	err = db.Get(&data, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.CommunityDetail{}, errors.New("数据为空")
		}
		return
	}
	return
}
