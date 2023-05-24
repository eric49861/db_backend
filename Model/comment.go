package model

type Comment struct {
	Id         uint   `gorm:"primaryKey;column:id"`
	CreatorId  uint   `gorm:"not null;column:creator_id"`
	QuestionId uint   `gorm:"not null;column:question_id"`
	CommentId  uint   `gorm:"not null;column:comment_id"`
	Content    string `gorm:"not null;column:content"`
	Likes      uint   `gorm:"not null;column:likes"`
	CreateTime int64  `gorm:"autoCreateTime:nano;column:create_time"`
}
