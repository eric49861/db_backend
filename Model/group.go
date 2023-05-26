package model

type Group struct {
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	CreatorId  uint   `json:"creatorId" gorm:"not null;column:creator_id"`
	GroupName  string `json:"groupName" gorm:"not null;column:group_name"`
	Topic      string `json:"topic" gorm:"not null;column:topic"`
	Notice     string `json:"notice" gorm:"not null;column:notice"`
	CreateTime int64  `json:"createTime" gorm:"not null;autoCreateTime:nano;column:create_time"`
	EndTime    int64  `json:"endTime" gorm:"not null;column:end_time"`
	State      bool   `json:"state" gorm:"not null;column:state"`
	Remarks    string `json:"remarks" gorm:"not null;column:remarks"`
}

func (Group) TableName() string {
	return "group"
}
