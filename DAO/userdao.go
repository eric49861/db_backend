package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	db "studygroup-gin/Database/mysql"
	model "studygroup-gin/Model"
)

// GetUserById 通过用户的ID查询用户
func GetUserById(id uint64) (model.User, error) {
	user := model.User{}
	res := db.DBsql.First(&user, id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return user, fmt.Errorf("数据库中不存在 id=%d 的用户", id)
	}
	return user, nil
}

// GetUserByUsername 通过用户名查询用户
func GetUserByUsername(username string) (model.User, error) {
	user := model.User{}
	res := db.DBsql.Where("username = ?", username).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return user, fmt.Errorf("数据库中不存在 'username' 为 '%s' 的用户", username)
	} else {
		return user, nil
	}
}

// GetUserByUsernameAndPassword 通过用户名和密码查询用户
func GetUserByUsernameAndPassword(username string, password string) (model.User, error) {
	user := model.User{}
	res := db.DBsql.Where("username = ? AND password = ?", username, password).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return user, fmt.Errorf("数据库中不存在 'username' 为 '%s', 'password' 为 '%s' 的用户", username, password)
	} else {
		return user, nil
	}
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(user model.User) error {
	err := db.DBsql.Save(&user).Error
	return err
}
