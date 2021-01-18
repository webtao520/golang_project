package main

import (
	"book/db"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化数据库
	err := db.InitDB()
	if err != nil {
		panic(err)
	}

	// 创建路由
	// 默认使用了2个中间件logger() 和 recovery()
	r := gin.Default()
	//  加载页面
	r.LoadHTMLGlob("templates/*")
	// 查看所有图书
	r.GET("book/list", bookListHandler)
	// 添加图书页面
	r.GET("book/new", newBook)
	// 添加图书 后端逻辑
	r.POST("book/new", saveBook)
	// 删除图书
	r.GET("/book/delete", deleteBook)
	_ = r.Run(":8000")
}

func bookListHandler(c *gin.Context) {
	bookList, err := db.QueryAllBook()
	if err != nil {
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	} else {
		// fmt.Println("====>", bookList) [0xc00020b0a0 0xc00020b140 0xc00020b1e0]
		// for k, v := range bookList {
		// 	fmt.Println(k, v.Price)
		// }
		// 返回数据
		c.HTML(http.StatusOK, "book_list.html", gin.H{
			"code": 0,
			"data": bookList,
		})
	}
}

func newBook(c *gin.Context) {
	c.HTML(http.StatusOK, "new_book.html", gin.H{})
}

func saveBook(c *gin.Context) {
	//  获取前台表单数据
	title := c.PostForm("title")
	price := c.PostForm("price")
	priceInt64, _ := strconv.ParseInt(price, 10, 64)
	err := db.InsertBook(title, priceInt64)
	if err != nil {
		panic(err.Error())
	}
}

// 删除图书
func deleteBook(c *gin.Context) {

	id := c.DefaultQuery("id", "0")
	idInt64, _ := strconv.ParseInt(id, 10, 64)
	err := db.Delete(idInt64)
	if err != nil {
		fmt.Println("删除失败")
		return
	}
	return
}
