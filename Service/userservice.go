package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	dao "studygroup-gin/DAO"
	db "studygroup-gin/Database/mysql"
	model "studygroup-gin/Model"
	"studygroup-gin/util"
)

//和用户相关的服务放在此处

// DoLogin 执行登录逻辑
func DoLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		name := context.PostForm("username")
		password := context.PostForm("password")
		println("name:", name)
		println("password: ", password)
		// 从数据库中查询该用户是否存在，如果存在，则生成一个token返回给前端
		user := model.User{}
		result := db.DBsql.Where("username = ? AND password = ?", name, util.EncryptWithMD5(password)).First(&user)
		// 没有找到该用户
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "用户名或者密码错误",
			})
			return
		}
		// 根据用户信息生成token并返回
		tokenString, err := util.GenerateTokenWithHS256(user.UserName, user.Id)
		if err != nil {
			panic(err)
		} else {
			//在响应头中添加token字段
			context.Header("Token", tokenString)
			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "登录成功",
				"user": user,
			})
		}
	}
}

// DoSignup 执行注册逻辑
func DoSignup() gin.HandlerFunc {
	return func(context *gin.Context) {
		name := context.PostForm("username")
		result := db.DBsql.Where("username = ?", name).First(&model.User{})
		//用户名已被占用
		if result.RowsAffected != 0 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "用户名已被占用",
			})
			return
		}
		user := model.User{
			UserName: context.PostForm("username"),
			Password: util.EncryptWithMD5(context.PostForm("password")),
			Gender:   context.PostForm("gender"),
			QQ:       context.PostForm("qq"),
		}
		result = db.DBsql.Create(&user)
		if result.Error == nil {
			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "注册成功",
			})
		} else {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "发生为止错误",
			})
		}
	}
}

// ModifyUserInfo 更新用户的个人信息
func ModifyUserInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取前端提交的表单信息并封装成对象
		id, _ := strconv.ParseUint(context.PostForm("id"), 10, 64)
		username := context.PostForm("username")
		age, _ := strconv.ParseUint(context.PostForm("age"), 10, 64)
		gender := context.PostForm("gender")
		occupation := context.PostForm("occupation")
		signature := context.PostForm("signature")

		user, _ := dao.GetUserById(id)
		user.UserName = username
		user.Age = age
		user.Gender = gender
		user.Occupation = occupation
		user.Signature = signature

		//调用userdao中的方法更新用户
		err := dao.UpdateUserInfo(user)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "用户信息更新失败",
			})
			return
		}
		newUser, _ := dao.GetUserById(id)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "用户信息更新成功",
			"user": newUser,
		})
	}
}
