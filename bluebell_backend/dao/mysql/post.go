package mysql

import "bluebell_backend/models"

func CreatePost(p *models.Post) error {
	sqlStr := "insert into post(post_id,title,content,author_id,community_id) values(?,?,?,?,?)"
	_, err := db.Exec(sqlStr, p.PostID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return err
}

func GetPostById(id int64) (data *models.Post, err error) {
	data = new(models.Post)
	sqlStr := "select post_id,content,title,author_id,community_id,create_time from post where post_id=?"
	err = db.Get(data, sqlStr, id)
	return
}
