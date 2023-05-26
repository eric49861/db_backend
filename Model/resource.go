package model

type Resource struct {
	Id         uint   `json:"id" gorm:"primaryKey;column:id"`
	UserId     uint   `json:"user_id" gorm:"column:user_id"`
	GroupId    uint   `json:"group_id" gorm:"not null;column:group_id"`
	Content    string `json:"content" gorm:"not null;column:content"`
	UploadTime int64  `json:"upload_time" gorm:"not null;column:upload_time"`
	UpdateTime int64  `json:"update_time" gorm:"autoCreateTime:nano;column:update_time"`
}

func (Resource) TableName() string {
	return "resource"
}
