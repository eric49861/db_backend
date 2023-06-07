package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	util "studygroup-gin/util"
)

func ValidToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.PostForm("token")
		fmt.Println("tokenStirng: ", tokenString)
		if util.IsValidTokenString(tokenString) {
			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "success",
			})
		} else {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "非法token",
			})
		}
	}
}
