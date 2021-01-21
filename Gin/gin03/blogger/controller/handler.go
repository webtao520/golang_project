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

	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": allCategoryList,
	})

}

// 创建文章
func NewArticle(c *gin.Context) {
	// 获取所有文章分类
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get article failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	c.HTML(http.StatusOK, "views/post_article.html", categoryList) // 根节点
}

// 保存文章
func ArticleSubmit(c *gin.Context) {
	author := c.PostForm("author")
	title := c.PostForm("title")
	categoryIdStr := c.PostForm("category_id")
	content := c.PostForm("content")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	//保存文章数据
	err = service.InsertArticle(content, author, title, categoryId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

// 文章详情
func ArticleDetail(c *gin.Context) {
	articleIdStr := c.Query("article_id") //get 方式获取文章id号
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 获取文章详情
	articleDetail, err := service.GetArticleDetail(articleId)
	if err != nil {
		fmt.Printf("get article detail failed,article_id:%d err:%v\n", articleId, err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	fmt.Println(articleDetail)
}
