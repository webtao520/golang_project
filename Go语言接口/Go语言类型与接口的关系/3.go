package main 

import (
	"fmt"
)

//  一个服务需要满足能够开启和写日志的功能
type  Service interface {
	  Start () // 开启服务
	  Log(string) // 日志输出
}

 // 日志器
 type Logger  struct {
	 
 }

 // 实现 Service 的Log () 方法
 func (g *Logger)Log(l string){
	  fmt.Println(l)
 }

 // 游戏服务
 type GameService  struct {
	Logger  // 嵌入日志器
 }


 // 实现Service的Start()方法
func (g *GameService) Start(){
	fmt.Println("Start")
}

func main(){
   var s Service = new (GameService)
   s.Start()
   s.Log("golang")
   // s 就可以使用 Start() 方法和 Log() 方法，其中，Start() 由 GameService 实现，Log() 方法由 Logger 实现。
}


/**
PS D:\goLang\github\golang_project\Go语言接口\Go语言类型与接口的关系> go run 3.go
Start
golang

代码说明如下：
定义服务接口，一个服务需要实现 Start() 方法和日志方法。
定义能输出日志的日志器结构。
为 Logger 添加 Log() 方法，同时实现 Service 的 Log() 方法。
定义 GameService 结构。
在 GameService 中嵌入 Logger 日志器，以实现日志功能。
GameService 的 Start() 方法实现了 Service 的 Start() 方法。

*/