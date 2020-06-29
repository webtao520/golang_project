
# golang  语言 研究

### 细说通道类型

   > 通道类型篇（一）：基本语法和缓冲通道

   > 通道类型篇（二）：单向通道及其使用

   > 通道类型篇（三）：通过 select 语句等待通道就绪
   
   > 通道类型篇（四）：错误和异常处理


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

