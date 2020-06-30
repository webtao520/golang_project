## 先举个例子：

> 在 Go http包的Server中，每一个请求在都有一个对应的 goroutine 去处理。请求处理函数通常会启动额外的 goroutine 用来访问后端服务，比如数据库和RPC服务。用来处理一个请求的 goroutine 通常需要访问一些与请求特定的数据，比如终端用户的身份认证信息、验证相关的token、请求的截止时间。 当一个请求被取消或超时时，所有用来处理该请求的 goroutine 都应该迅速退出，然后系统才能释放这些 goroutine 占用的资源。

对于多并发的情况下，传统的方案：等待组sync.WaitGroup以及通过通道channel的方式的问题就会显现出来；

对于等待组控制多并发的情况：只有所有的goroutine都结束了才算结束，只要有一个goroutine没有结束， 那么就会一直等，这显然对资源的释放是缓慢的；

而对于通道Channel的方式下：通过在main goroutine中像chan中发送关闭停止指令，并配合select，从而达到关闭goroutine的目的，这种方式显然比等待组优雅的多，但是在goroutine中在嵌套goroutine的情况就变得异常复杂。

##### 等待组例子：
 > 见 context.go 
