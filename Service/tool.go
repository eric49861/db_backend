package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"studygroup-gin/util"
)

func ValidToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Query("token")
		fmt.Println("tokenStirng: ", tokenString)
		if util.IsValidTokenString(tokenString) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "success",
			})
		} else {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "非法token",
			})
		}
	}
}
