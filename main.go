package main

import (
	"github.com/gin-gonic/gin"
	middle "studygroup-gin/Middleware"
	router "studygroup-gin/Router"
)

func main() {

	r := gin.Default()
	//跨域中间件
	r.Use(middle.Cors())
	//注册路由
	router.RegisteRouter(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8888") // 监听并在 0.0.0.0:8080 上启动服务
}
