package controller

import (
	"blogger/service"
	"fmt"
	_ "fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexHandle(c *gin.Context) {
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": allCategoryList,
	})
}

func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	articleRecordList, err := service.GetArticleRecordListById(int(categoryId), 0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	fmt.Println(articleRecordList)
	// fmt.Printf("%T\n", categoryId)
}
