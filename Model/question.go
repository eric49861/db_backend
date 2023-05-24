package model

type Question struct {
	Id         uint   `gorm:"primaryKey;column:id"`
	CreatorId  uint   `gorm:"not null;column:creator_id"`
	GroupId    uint   `gorm:"not null;column:group_id"`
	Likes      uint   `gorm:"not null;column:likes"`
	Content    string `gorm:"not null;column:content"`
	CreateTime int64  `gorm:"autoCreateTime:nano;column:create_time"`
}
