package dao_test

import (
	"fmt"
	dao "studygroup-gin/DAO"
	"testing"
)

func TestGetAllPost(t *testing.T) {
	posts := dao.GetAllPost()
	for index, v := range posts {
		fmt.Printf("index: %d, post: %v", index, v)
	}
}

func TestGetPostById(t *testing.T) {
	post := dao.GetPostById(1)
	fmt.Printf("%v", post)
}
