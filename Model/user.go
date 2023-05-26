package model

type User struct {
	Id         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"username" gorm:"not null;column:username"`
	Password   string `json:"password" gorm:"not null;column:password"`
	Gender     string `json:"gender" gorm:"not null;column:gender"`
	QQ         string `json:"qq" gorm:"not null;column:qq"`
	Integral   uint   `json:"integral" gorm:"not null;column:integral"`
	CreateTime int64  `json:"createTime" gorm:"autoCreateTime:nano;column:create_time"`
	UpdateTime int64  `json:"updateTime" gorm:"autoCreateTime:nano;column:create_time"`
}

func (User) TableName() string {
	return "user"
}
