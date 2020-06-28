package   main 

/*
   sync 包 提供了基本的同步基元，如互斥锁，除了once 和 waitGroup类型， 大部分都是适用于低水平程序线程，高水平的同步使用channel通信更好一些。
   本包的类型的值不应被拷贝
*/
import (
	"fmt"
	"sync"
)
// once  是只执行一次动作的对象
func main (){
   var once sync.Once
   onceBody :=func(){
	   fmt.Println("Only once")
   }
   done:=make(chan bool)
   for i:=0;i<10;i++{
	   go func(){
		   once.Do(onceBody)   // Do 方法当且仅当第一次被调用式才执行函数，换句话说，给定变量
		   done <- true
	   }()
   }
   
   for i:=0;i<10;i++ {
	   data:=<-done
	   fmt.Println("data====>",data)
   }
}