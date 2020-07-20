
# golang  语言学习笔记  [基础不牢固，地动山摇~]

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

### 爬虫案例 

 + 爬虫案例 CODE 

   * reptile
  
 + web网站数据抓取

   * golang 实现数据网站抓取项目

### 包的研究和学习

+ package

   * MySQL：github.com/go-sql-driver/mysql  常用的数据库orm
   * sync   RWMutex（读写锁），Mutex 互斥锁，sync.WaitGroup 类型
   * viper 热加载配置文件
   * net  熟悉http/https
   * select  chan 通道分支
   * singleton_pattern_config  sync 通过单例模式初始化全局配置，函数执行一次（优化性能）

### 语言学习路线图

### go 实现的案例demo
  
   * 用Go语言计算一个人的年龄，生肖，星座

### go应用
   * 定时器
   * 类型断言
   * 内置函数
   * go控制并发两种方式
   * goroutine 和 channel
   * 反射

### ETCD

+ ETCD

### go常见错误总结

   * 创建的链接，使用完要关闭
   * map 处理完map元素的任务，要删掉该map的元素，避免重复处理

### go面向对象
  
   - 接口
     * 接口实现
     * 接口多态
     * 接口继承
     * 接口转换
     * 空接口

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

### 秒杀项目

+ secKill

### Web 编程

+ chitchat 二次开发案例

   * 声明：本项目基于 Go Web Programming 一书中的 chitchat 做的二次开发，在原项目基础上将数据库调整为 MySQL、路由器调整为了 gorilla/mux、调整了整体目录结构、新增了配置文件单例模式获取、本地化编程以及应用部署流程。   

### Web 框架

### 插件/package

   *  发送邮件库 https://github.com/go-gomail/gomail
   * 读写Microsoft Excel https://github.com/360EntSecGroup-Skylar/excelize 详细资料
   * 生成uuid https://github.com/satori/go.uuid
   * 开源，分布式，简单高效的搜索引擎 https://github.com/go-ego/riot
   * 基于 Go 的高性能 MySQL Proxy https://github.com/flike/kingshard
   * yaml对Go语言的支持 https://github.com/go-yaml/yaml/tree/v2.2.7
   * Codis是一个分布式Redis解决方案数据库代理 https://github.com/CodisLabs/codis
   * 用Go语言编写的markdown解析器 https://github.com/yuin/goldmark

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

### ORM 引擎文档 
   * https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-01/index.html
