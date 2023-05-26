package router

import (
	"github.com/gin-gonic/gin"
	middle "studygroup-gin/Middleware"
	service "studygroup-gin/Service"
)

func RegisteRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		free := v1.Group("/1")
		{
			free.POST("/login", service.DoLogin())
			free.POST("/signup", service.DoSignup())
		}
		//该类接口需要token验证
		authentication := v1.Group("/2")
		{
			authentication.Use(middle.JWTAuthMiddleware())
			authentication.GET("/get-all-group", service.GetGroupList())
		}

		// 其他的一些工具接口
		tool := v1.Group("/3")
		{
			tool.POST("/valid-token", service.ValidToken())
		}
	}
}
