package db

import (
	 "fmt"
	"testing"
	_"github.com/jmoiron/sqlx"
)

// 数据库连接初始化
// parseTime=true 将mysql中时间类型(timestamp)，自动解析为go结构体中的时间类型 time.Time
// 不加报错
func init() {
	addr := "root:root123@tcp(localhost:3306)/test?parseTime=true"
	err:=Init(addr)
	if err !=nil {
		panic(err)
	}
	return
}

// 测试获取相关文章
func TestGetRelativeAricleList(t *testing.T) {
	  var  articleId int64
	  articleId = 1
	  articleList,err:=GetRelativeArticle(articleId)
	if err !=nil {
		t.Errorf("get category failed, err:%v\n",err)
	}
	for _,v:=range articleList{
		//t.Logf(v)
		fmt.Println(*v) // * 取值
	}
}


func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := GetAllCategoryList()
	if err != nil {
		t.Errorf("get category  failed, err:%v\n", err)
		return
	}
	for _, v := range categoryList {
		t.Logf("category:%#v", v)
	}
}