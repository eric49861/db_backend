package middle

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
	db "studygroup-gin/Database/mysql"
	model "studygroup-gin/Model"
	"studygroup-gin/util"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//检查请求头中是否携带了token
		tokenString := context.Request.Header.Get("Token")
		//如果没有携带，则返回未认证错误
		if strings.Compare(tokenString, "") == 0 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "未认证错误",
			})
			return
		}
		//携带了token，调用验证token的方法对token进行验证
		id, err := util.ParseToken(tokenString)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "无效的令牌",
			})
			return
		}
		// 否则说明token是有效的，刷新token并将新的token放入响应头中
		newTokenString, err := util.RefreshToken(tokenString)
		if err != nil {
			// 刷新失败
		} else {
			context.Header("new-token", newTokenString)
		}
		// 从数据库中查询该用户，看是否存在
		var user model.User
		result := db.DBsql.First(&user, id)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "用户不存在",
			})
			return
		}
		//将该用户的对象添加到上下文中
		context.Set("user", user)
		context.Next()
	}
}
