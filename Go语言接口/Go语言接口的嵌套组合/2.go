package main

import (
	"io"
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

// 声明一个设备结构
type device struct {
}

// 实现io.Writer的Write()方法
func (d *device) Write(p []byte)(n int,err error) {
	 return 0,nil
}

// 实现io.Closer的Close 方法
func (d *device)Close()error{
	return nil 
}

func main(){
	// 声明写入关闭器，并赋予device 的实例
	var wc io.WriteCloser = new(device)
	// 写入数据
	wc.Write(nil)
	//关闭设备
	wc.Close()
    // 声明写入器, 并赋予device的新实例
	var writeOnly io.Writer = new(device)
	
	// 写入数据
	writeOnly.Write(nil)
}


/**

代码说明如下：
第 8 行定义了 device 结构体，用来模拟一个虚拟设备，这个结构会实现前面提到的 3 种接口。
第 12 行，实现了 io.Writer 的 Write() 方法。
第 17 行，实现了 io.Closer 的 Close() 方法。
第 24 行，对 device 实例化，由于 device 实现了 io.WriteCloser 的所有嵌入接口，因此 device 指针就会被隐式转换为 io.WriteCloser 接口。
第 27 行，调用了 wc（io.WriteCloser接口）的 Write() 方法，由于 wc 被赋值 *device，因此最终会调用 device 的 Write() 方法。
第 30 行，与 27 行类似，最终调用 device 的 Close() 方法。
第 33 行，再次创建一个 device 的实例，writeOnly 是一个 io.Writer 接口，这个接口只有 Write() 方法。
第 36 行，writeOnly 只能调用 Write() 方法，没有 Close() 方法。

*/