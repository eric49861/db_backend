package dao

import (
	db "studygroup-gin/Database/mysql"
	model "studygroup-gin/Model"
)

func GetAllPost() ([]model.Post, error) {
	posts := make([]model.Post, 0)
	//查询所有的贴子
	res := db.DBsql.Find(&posts)
	return posts, res.Error
}

func GetPostById(id uint64) (model.Post, error) {
	post := model.Post{}
	res := db.DBsql.Where("id = ?", id).First(&post)
	return post, res.Error
}

// GetCommentsByPostId 根据帖子的id获取所有的评论
func GetCommentsByPostId(id uint64) ([]model.PostComment, error) {
	comments := make([]model.PostComment, 0)
	res := db.DBsql.Where("post_id = ?", id).Find(&comments)
	return comments, res.Error
}

func SaveComment(comment *model.PostComment) error {
	res := db.DBsql.Create(comment)
	return res.Error
}
