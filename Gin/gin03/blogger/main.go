package main

import (
	"blogger/controller"
	"blogger/dao/db"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func main() {
	//  1.定义实体
	// 2.数据层 数据库初始化
	err := db.Init()
	if err != nil {
		panic(err.Error())
	}
	// 创建路由
	router := gin.Default()
	ginpprof.Wrapper(router)
	router.Static("/static", "./static")
	// 加载模板文件
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandle)                   // 首页列表
	router.GET("/category", controller.CategoryList)          // 获取分类信息
	router.GET("/article/new/", controller.NewArticle)        // 创建文章视图层
	router.POST("/article/submit/", controller.ArticleSubmit) // 保存文章内容
	router.GET("/article/detail/", controller.ArticleDetail)  //文章详情

	_ = router.Run(":8008")
}
