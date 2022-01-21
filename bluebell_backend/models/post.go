package models

import "time"

type Post struct {
	PostID      int64     `json:"id" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id"`
	Status      int       `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 自定义帖子详情接口结构体
type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	// 嵌入其他结构体
	*Post
	*Community
}
