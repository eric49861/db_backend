package model

type Resource struct {
	Id         uint   `gorm:"primaryKey;column:id"`
	UserId     uint   `gorm:"column:user_id"`
	GroupId    uint   `gorm:"not null;column:group_id"`
	Content    string `gorm:"not null;column:content"`
	UploadTime int64  `gorm:"not null;column:upload_time"`
	UpdateTime int64  `gorm:"autoCreateTime:nano;column:update_time"`
}
