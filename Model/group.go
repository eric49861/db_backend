package model

type Group struct {
	Id         uint   `gorm:"primaryKey;autoIncrement;column:id"`
	CreatorId  uint   `gorm:"not null;column:creator_id"`
	GroupName  string `gorm:"not null;column:group_name"`
	Topic      string `gorm:"not null;column:topic"`
	Notice     string `gorm:"not null;column:notice"`
	CreateTime int64  `gorm:"not null;autoCreateTime:nano;column:create_time"`
	EndTime    int64  `gorm:"not null;column:end_time"`
	State      bool   `gorm:"not null;column:state"`
	Remarks    string `gorm:"not null;column:notice"`
}
