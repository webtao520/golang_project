package main 

import (
	"fmt"
	"sync/atomic"
)

var (
	// 序列号
	seq int64 
)

// 序列号生成器
func GenID  () int64 {
	// 尝试原子的增加序列号
	// 根据 atomic.AddInt64() 的参数声明，这个函数会将修改后的值以返回值方式传出。下面代码对加粗部分进行了修改：
	return   atomic.AddInt64(&seq, 1)
    //return seq
}

func main(){
 //生成10个并发序列号
 for i := 0; i < 10; i++ {
	go GenID()
}
fmt.Println(GenID())
}