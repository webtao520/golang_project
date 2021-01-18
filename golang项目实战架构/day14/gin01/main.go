package main 


import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin 的helloworld

func  main(){
	// 1. 创建路由
r:=gin.Default()
// 2. 绑定路由规则，执行的函数
 // gin.Context，封装了request和response
 r.Get("/",func(c *gin.Context){
    c.String(http.StatusOK, "hello World!")
 })
 // 3.监听端口，默认在8080
 r.Run(":8000")
}