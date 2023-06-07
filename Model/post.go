package model

type Post struct {
	Id         uint64 `json:"id" gorm:"primaryKey;column:id"`
	CreatorId  uint64 `json:"creator_id" gorm:"not null;column:creator_id"`
	Likes      uint64 `json:"likes" gorm:"not null;column:likes"`
	Content    string `json:"content" gorm:"not null;column:content"`
	CreateTime int64  `json:"create_time" gorm:"autoCreateTime:nano;column:create_time"`
}

func (Post) TableName() string {
	return "post"
}
