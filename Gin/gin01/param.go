package main 


import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	// 1. 创建路由
	// 默认使用了 2 个中间件logger(), Recovery()
	r:=gin.Default()
	// api 参数
	r.Get("/user/:name/*action", func(c *gin.Context) {
		 // 
		 name:=c.Params("name")
		 action:=c.Params("action")
		 c.String(http.StatusOK,name+ " is "+action)
	})
	r.Run(":8080")
}