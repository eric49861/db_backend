package model

type Question struct {
	Id         uint   `json:"id" gorm:"primaryKey;column:id"`
	CreatorId  uint   `json:"creator_id" gorm:"not null;column:creator_id"`
	GroupId    uint   `json:"group_id" gorm:"not null;column:group_id"`
	Likes      uint   `json:"likes" gorm:"not null;column:likes"`
	Content    string `json:"content" gorm:"not null;column:content"`
	CreateTime int64  `json:"create_time" gorm:"autoCreateTime:nano;column:create_time"`
}

func (Question) TableName() string {
	return "question"
}
