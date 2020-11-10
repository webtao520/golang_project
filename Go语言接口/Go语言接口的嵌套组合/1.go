package main 


import (
	"fmt"
)

type Writer interface {
	 Write(p []byte)(n int,err error)
}

type Closer interface {
	Close() error
} 

type WriteCloser interface {
	Writer
	Closer
}


func main(){

}

/**

代码说明如下：
第 1 行定义了写入器（Writer），如这个接口较为常用，常用于 I/O 设备的数据写入。
第 5 行定义了关闭器（Closer），如有非托管内存资源的对象，需要用关闭的方法来实现资源释放。
第 9 行定义了写入关闭器（WriteCloser），这个接口由 Writer 和 Closer 两个接口嵌入。也就是说，WriteCloser 同时拥有了 Writer 和 Closer 的特性。

*/