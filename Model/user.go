package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id         uint64  `json:"id" gorm:"primaryKey"`
	UserName   string  `json:"username" gorm:"not null;column:username"`
	Password   string  `json:"password" gorm:"not null;column:password"`
	Signature  string  `json:"signature" gorm:"not null;column:signature"`
	Avatar     string  `json:"avatar" gorm:"not null;column:avatar"`
	Gender     string  `json:"gender" gorm:"not null;column:gender"`
	Age        uint64  `json:"age" gorm:"not null;column:age"`
	QQ         string  `json:"qq" gorm:"not null;column:qq"`
	Integral   uint64  `json:"integral" gorm:"not null;column:integral;default:0"`
	Occupation string  `json:"occupation" gorm:"not null;column:occupation"`
	StudyTime  float64 `json:"studyTime" gorm:"not null;column:study_time"`
	CreateTime int64   `json:"createTime" gorm:"autoCreateTime:nano;column:create_time"`
	UpdateTime int64   `json:"updateTime" gorm:"autoCreateTime:nano;column:update_time"`
}

func (User) TableName() string {
	return "user"
}

func (self *User) BeforeSave(tx *gorm.DB) error {
	self.UpdateTime = time.Now().UnixNano()
	return nil
}
