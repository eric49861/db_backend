package model

type Comment struct {
	Id         uint64 `json:"id" gorm:"primaryKey;column:id"`
	CreatorId  uint64 `json:"creator_id" gorm:"not null;column:creator_id"`
	QuestionId uint64 `json:"question_id" gorm:"not null;column:question_id"`
	Content    string `json:"content" gorm:"not null;column:content"`
	Likes      uint64 `json:"likes" gorm:"not null;column:likes"`
	CreateTime int64  `json:"create_time" gorm:"autoCreateTime:nano;column:create_time"`
}

func (Comment) TableName() string {
	return "comment"
}
