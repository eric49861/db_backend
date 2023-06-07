package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	dao "studygroup-gin/DAO"
	model "studygroup-gin/Model"
)

type ZoneResult struct {
	Post model.Post
	User model.User
}

func GetAllPost() gin.HandlerFunc {
	return func(context *gin.Context) {
		//查询数据库，获取所有的贴子
		posts, err := dao.GetAllPost()
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "bad server",
			})
			return
		}
		result := make([]ZoneResult, 0)
		for _, post := range posts {
			user, err := dao.GetUserById(post.CreatorId)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"msg": "some error happened",
				})
				panic(err)
			}
			result = append(result, ZoneResult{
				Post: post,
				User: user,
			})
		}

		//返回数据
		context.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"result": result,
		})
	}
}

func GetUserById() gin.HandlerFunc {
	return func(context *gin.Context) {
		ids := context.Query("id")
		id, err := strconv.ParseUint(ids, 10, 64)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "无效的参数",
			})
		}
		user, err := dao.GetUserById(id)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}

func GetCommentsByPostId() gin.HandlerFunc {
	return func(context *gin.Context) {
		ids := context.Query("id")
		id, err := strconv.ParseUint(ids, 10, 64)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "无效的参数",
			})
			return
		}
		comments, err := dao.GetCommentsByPostId(id)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "bad server",
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"comments": comments,
		})
	}
}

func SaveComment() gin.HandlerFunc {
	return func(context *gin.Context) {
		content := context.PostForm("content")
		postIds := context.PostForm("postId")
		ids := context.PostForm("id")
		id, err := strconv.ParseUint(ids, 10, 64)
		postId, err := strconv.ParseUint(postIds, 10, 64)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "无效的参数",
			})
		}
		comment := model.PostComment{
			Id:        0,
			CreatorId: id,
			PostId:    postId,
			Content:   content,
			Likes:     0,
		}
		err = dao.SaveComment(&comment)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "bad server",
			})
		}
		context.JSON(http.StatusOK, gin.H{
			"msg": "评论成功",
		})
	}
}
