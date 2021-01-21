package db

import (
	"fmt"
	"testing"
)

func init() {
	addr := "root:root123@tcp(localhost:3306)/test?parseTime=true"
	err:=Init(addr)
	if err !=nil {
		panic(err)
	}
	return
}

// 根据当前id 获取上一篇文章
func TestGetPrevArticleById(t *testing.T) {
		info,err:=GetPrevArticleById(3)
		if err !=nil {
			panic(err)
		}
		fmt.Println(info)
}

// 根据当前的文章id 获取下一篇文章
func  TestGetNextArticleById(t  *testing.T){
	info,err := GetNextArticleById(2)
	if err !=nil {
		panic(err)
	}
	fmt.Println(info)
}