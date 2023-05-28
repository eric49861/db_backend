package model

type User struct {
	Id         uint    `json:"id" gorm:"primaryKey"`
	Name       string  `json:"username" gorm:"not null;column:username"`
	Password   string  `json:"password" gorm:"not null;column:password"`
	Signature  string  `json:"signature" gorm:"not null;column:signature"`
	Avatar     string  `json:"avatar" gorm:"not null;column:avatar"`
	Gender     string  `json:"gender" gorm:"not null;column:gender"`
	QQ         string  `json:"qq" gorm:"not null;column:qq"`
	Integral   uint    `json:"integral" gorm:"not null;column:integral;default:0"`
	Occupation string  `json:"occupation" gorm:"not null;column:occupation"`
	StudyTime  float64 `json:"studyTime" gorm:"not null;column:study_time"`
	CreateTime int64   `json:"createTime" gorm:"autoCreateTime:nano;column:create_time"`
	UpdateTime int64   `json:"updateTime" gorm:"autoCreateTime:nano;column:create_time;default:"`
}

func (User) TableName() string {
	return "user"
}
