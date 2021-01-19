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
	router.GET("/", controller.IndexHandle)
	_ = router.Run(":8008")
}
