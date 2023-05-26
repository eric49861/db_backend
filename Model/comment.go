package model

type Comment struct {
	Id         uint   `json:"id" gorm:"primaryKey;column:id"`
	CreatorId  uint   `json:"creator_id" gorm:"not null;column:creator_id"`
	QuestionId uint   `json:"question_id" gorm:"not null;column:question_id"`
	CommentId  uint   `json:"comment_id" gorm:"not null;column:comment_id"`
	Content    string `json:"content" gorm:"not null;column:content"`
	Likes      uint   `json:"likes" gorm:"not null;column:likes"`
	CreateTime int64  `json:"create_time" gorm:"autoCreateTime:nano;column:create_time"`
}

func (Comment) TableName() string {
	return "comment"
}
