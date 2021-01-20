
# golang  语言学习笔记  [基础才最重要。。 要多学一下~~]

### 并发性Concurrency

### 错误和异常处理

   * error 接口及其使用
   * defer 语句及使用示例
   * panic 和 recover

### 数据库ORM

### 为什么需要context包

### 细说通道类型

   * 通道类型篇（一）：基本语法和缓冲通道
   * 通道类型篇（二）：单向通道及其使用
   * 通道类型篇（三）：通过 select 语句等待通道就绪
   * 通道类型篇（四）：错误和异常处理

### golang项目实战架构基础与实战

### ⽇志收集项⽬架构设计及Kafka介绍

### 细说go语言数据类型

   * 指针
   * go语言中byte类型和rune类型


### 字符串中常用的系统函数

   * 10进制转换
   * 整数转字符串
   * 字符串转整数
   * 字符串转byte数组
   * len 获取字符串长度
   * rune 类型
   * ...

### 单元测试

### golang_project  [golang开源项目适合练手]

   * chitchat 博客项目
   * reptile  爬虫案例
   * golang实现数据网站抓取项目
   * newSecKill go新秒杀项目
   * secKill  go秒杀项目
   * division_provinces 省份划分
  

### 包的研究和学习

+ package

   * MySQL：github.com/go-sql-driver/mysql  常用的数据库orm
   * sync   RWMutex（读写锁），Mutex 互斥锁，sync.WaitGroup 类型
   * viper 热加载配置文件
   * net  熟悉http/https
   * select  chan 通道分支
   * singleton_pattern_config  sync 通过单例模式初始化全局配置，函数执行一次（优化性能）
   * sort
   * strings
   * sync
   * time
   * viper
   * md5
   * context
   * mysql
   * atomic 「atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用」
   * fmt  「mt包实现了类似C语言printf和scanf的格式化I/O。格式化动作（'verb'）源自C语言但更简单」
   * url  「url包解析URL并实现了查询的逸码」
   * bytes 「bytes包实现了操作[]byte的常用函数。本包的函数和strings包的函数相当类似」
   * strconv 「strconv 包实现了基本数据类型和其字符串表示的相互转换」 

### 语言学习路线图

### go 实现的案例demo
  
   * 用Go语言计算一个人的年龄，生肖，星座

### go应用
   * 定时器
   * 类型断言
   * 内置函数
   * 锁
   * go控制并发两种方式
   * goroutine 和 channel
   * 反射
   * slice
   * map
   * slice切片去重(传入的是string类型,传入的是int类型)

### Go语言包（package）
   * 包的基本概念
   * Go语言封装简介及实现细节   


### Go语言并发  
   * 并发和并行的区别
   * Go语言单向通道
   * Go语言单向通道
   * Go语言轻量级线程
   * Go语言通道（chan）
   * Go语言无缓冲的通道
   * Go语言带缓冲的通道
   * Go语言channel超时机制
   * Go语言通道的多路复用
   * 互斥锁和读写互斥锁
   * 检测代码在并发环境下可能出现的问题
   * 示例：并发打印
   * 示例：使用通道响应计时器的事件
   * Go语言等待组
   * Go语言多核并行化
   * Go语言关闭通道后继续使用通道
   * Go语言CSP：通信顺序进程简述
   * 使用select切换协程

### Go语言反射

### Gin 框架

### ETCD

+ ETCD

### go常见错误总结

   * go语言坑之for range
   * 创建的链接，使用完要关闭
   * map 处理完map元素的任务，要删掉该map的元素，避免重复处理

### go面向对象
  
   - 接口
     * 接口实现
     * 接口多态
     * 接口继承
     * 接口转换
     * 空接口

### Go容器
   * Go语言多维数组
   * Go语言数组
   * Go语言切片
   * 使用append()为切片添加元素  
   * Go语言遍历map
   * Go语言从切片中删除元素
   * Go语言切片复制
   * Go语言list（列表）
   * Go语言make和new关键字的区别及实现原理
   * Go语言map（映射）
   * Go语言map的多键索引
   * Go语言nil：空值
   * Go语言range关键字
   * Go语言sync.Map
   * map元素的删除和清空
   * map元素的删除和清空

 ### 语言接口
   * 部署Go语言程序到Linux服务器
   * 接口的nil判断
   * 示例：表达式求值器
   * 示例：使用空接口实现可以保存任意值的字典
   * Go语言接口的嵌套组合
   * Go语言接口和类型之间的转换
   * Go语言接口内部实现
   * Go语言接口声明（定义）
   * Go语言空接口类型
   * Go语言类型断言
   * Go语言类型分支（switch判断空接口中变量的类型）
   * Go语言类型与接口的关系
   * Go语言排序
   * Go语言实现接口的条件
   * Go语言error接口

### Go语言结构体
   * 初始化结构体的成员变量
   * 初始化内嵌结构体
   * 结构体内嵌模拟类的继承
   * 类型内嵌和结构体内嵌
   * 示例：使用事件系统实现事件的响应和处理
   * 为任意类型添加方法
   * Go语言方法和接收器
   * Go语言构造函数
   * Go语言结构体定义
   * Go语言实例化结构体
   * 内嵌结构体成员名字冲突
   * Go语言垃圾回收和SetFinalizer
   * 示例：将结构体数据保存为JSON格式数据
   * Go语言链表操作

### Go语言文件处理
   * Go语言使用Gob传输数据
   * Go语言自定义数据文件
   * Go语言JSON文件的读写操作
   * Go语言XML文件的读写操作
   * Go语言使用Gob传输数据
   

### 设计模式

+ 创建型模式

   * 单例模式
   * 抽象工厂模式
   * 工厂方法模式
   * 生成器模式
   * 原型模式

+ 结构型模式

+ 行为型模式

### goroutine

+ goroutine

### json序列化与反序列化

+ json序列化与反序列化

### go uuid

+ uuid

### Web 编程

+ chitchat 二次开发案例

   * 声明：本项目基于 Go Web Programming 一书中的 chitchat 做的二次开发，在原项目基础上将数据库调整为 MySQL、路由器调整为了 gorilla/mux、调整了整体目录结构、新增了配置文件单例模式获取、本地化编程以及应用部署流程。   

### Web 框架
   * gin https://github.com/skyhee/gin-doc-cn#db-mysql

### 插件/package

   * 发送邮件库 https://github.com/go-gomail/gomail
   * 读写Microsoft Excel https://github.com/360EntSecGroup-Skylar/excelize 详细资料
   * 生成uuid https://github.com/satori/go.uuid
   * 开源，分布式，简单高效的搜索引擎 https://github.com/go-ego/riot
   * 基于 Go 的高性能 MySQL Proxy https://github.com/flike/kingshard
   * yaml对Go语言的支持 https://github.com/go-yaml/yaml/tree/v2.2.7
   * Codis是一个分布式Redis解决方案数据库代理 https://github.com/CodisLabs/codis
   * 用Go语言编写的markdown解析器 https://github.com/yuin/goldmark
   * 获取配置文件信息包 https://ini.unknwon.io/docs/intro/getting_started
   * mysql 第三方orm   https://github.com/jmoiron/sqlx
   * mysql 驱动     https://github.com/go-sql-driver/mysql
   * pprof 调试    	https://github.com/DeanThompson/ginpprof

### 项目

   * 基于beego框架的接口在线文档管理系统 https://github.com/lifei6671/mindoc
   * 开源文库系统 https://github.com/truthhun/DocHub
   * Go常用规范定义案例 https://github.com/zc2638/go-standard
   * 开源监控度量的看板系统 https://github.com/zc2638/go-standard
   * go应用开发的调试工具 https://github.com/derekparker/delve
   * 高并发、重量级爬虫软件 https://github.com/henrylee2cn/pholcus
   * Web分析 https://github.com/matomo-org/matomo

### 基于beego，jquery easyui ,bootstrap的一个后台管理系统

   * https://github.com/beego/admin

### Golang 设计模式文档

   * https://design-patterns.readthedocs.io/zh_CN/latest/creational_patterns/simple_factory.html#id2

### Golang 推荐博客
   * https://www.liwenzhou.com/posts/Go/go_menu/   

### ORM 引擎文档 
   * https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-01/index.html

### web编程实战
   * http://wen.topgoer.com/docs/golangweb/golangweb-1ck1klaugql7h   

