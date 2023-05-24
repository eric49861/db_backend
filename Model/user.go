package model

type User struct {
	Id         uint `gorm:"primaryKey"`
	Name       string
	Password   string
	QQ         string
	Integral   uint
	CreateTime int64
	UpdateTime int64
}
