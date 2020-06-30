
# golang  语言 研究 教程


### 细说通道类型

   1. 通道类型篇（一）：基本语法和缓冲通道
   2. 通道类型篇（二）：单向通道及其使用
   3. 通道类型篇（三）：通过 select 语句等待通道就绪
   4. 通道类型篇（四）：错误和异常处理


### 错误和异常处理

   1. error 接口及其使用
   2. defer 语句及使用示例
   3. panic 和 recover


### 爬虫案例 

 1. 爬虫案例 CODE 

   > reptile
  
 2. web网站数据抓取

   > golang 实现数据网站抓取项目

### go 包的研究和学习

   > package

      1. MySQL：github.com/go-sql-driver/mysql  常用的数据库orm
      2. sync   RWMutex（读写锁），Mutex 互斥锁，sync.WaitGroup 类型
      3. viper 热加载配置文件
      4. net  熟悉http/https
      5. select  chan 通道分支
      6. singleton_pattern_config  sync 通过单例模式初始化全局配置，函数执行一次（优化性能）

       
### go 并发 

   > goroutine

### go 并发理论

   > 并发性Concurrency

### Go Web 编程

   > chitchat 二次开发案例

   1.声明：本项目基于 Go Web Programming 一书中的 chitchat 做的二次开发，在原项目基础上将数据库调整为 MySQL、路由器调整为了 gorilla/mux、调整了整体目录结构、新增了配置文件单例模式获取、本地化编程以及应用部署流程。   

### Go Web 框架

| 项目 	| 简介 	| Star数 	|
|-	|-	|-	|
| gin-gonic/gin 	| Gin 是一个用 Go 语言开发的 Web 框架，提供类 Martini 的 API，但是性能更好。因为有了 httprouter 性能提升了 40 倍之多。 	| 36433 	|
| astaxie/beego 	| beego是一个用Go开发的应用框架，思路来自于tornado，路由设计来源于sinatra 	| 23525 	|
| kataras/iris 	| 通过Iris-Go，可以方便的帮助你来开发基于web的应用。简单来说：Iris-Go与国内大牛的BeeGo类似，但从其官方介绍的资料来看，Iris-Go的性能更优！ 	| 17776 	|
| labstack/echo 	| Echo 是个快速的 HTTP 路由器（零动态内存分配），也是 Go 的微型 Web 框架。 	| 16793 	|
| codegangsta/martini 	| Martini 是一个非常新的 Go 语言的 Web 框架，使用 Go 的 net/http 接口开发，类似 Sinatra 或者 Flask 之类的框架，你可使用自己的 DB 层、会话管理和模板 	| 10888 	|
| hoisie/web 	| web.go 跟 web.py 类似，但使用的是 Go 编程语言实现的 Web 应用开发框架。Go发布没多久该框架就诞生了，差不多是最早的Go框架。目前已经有段时间没有更新了。不过，该框架代码不多，其源码可以读一读。 	| 3440 	|
| go-macaron/macaron 	| Macaron 是一个具有高生产力和模块化设计的 Go Web 框架。框架秉承了 Martini 的基本思想，并在此基础上做出高级扩展 	| 2952 	|
| gernest/utron 	| utron 是一个 Go 语言轻量级的 MVC 框架，用于快速构建可伸缩以及可靠的数据库驱动的 Web 应用。 	| 2159 	|
| olahol/melody 	| Melody 是一个 Go 语言的微型 WebSocket 框架，基于 github.com/gorilla/websocket 开发， 	| 1771 	|
| henrylee2cn/faygo 	| Faygo 是一款快速、简洁的Go Web框架，可用极少的代码开发出高性能的Web应用程序（尤其是API接口）。只需定义 struct Handler，Faygo 就能自动绑定、验证请求参数并生成在线API文档。 	| 1503 	|
| lunny/tango 	| Tango，微内核可扩展的Go语言Web框架。同时支持函数和结构体作为执行体，插件丰富。 	| 835 	|
| robfig/revel 	| Revel 是 Go 语言的框架，其思路完全来自 Java 的 Play Framework。 	| 159 	|
| go-baa/baa 	| Baa 一个简单高效的Go web开发框架。主要有路由、中间件，依赖注入和HTTP上下文构成。 	|  	|   

