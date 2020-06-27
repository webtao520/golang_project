package main 


import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	   sync 同步包
		  同步：sync 
			 one by one 
			 串行
		  异步：async
			  同时执行
	 WaitGroup: 同步等待组
		 内置的计数器： 要执行的goroutine 的数量
		 Add(-), 设置计数器的数量
		 Done(), 将内置的计数器的数值减1，同Add(-1)
		 Wait() ,等待，导致执行wait 的goroutine进入阻塞状态，同步等待组中的计数器的值为0解除阻塞

	同步等待组：WaitGroup ,执行了wait的goroutine,要等待同步等待组中的其他的goroutine执行完毕	 
*/
func main (){

   var wg sync.WaitGroup
   //fmt.Printf("%T\n",wg)
   //fmt.Println(wg)
   wg.Add(2)
   //wg.Add(1)  
   go printNum1(&wg)
   go printNum2(&wg)
   wg.Wait()  // main 进入阻塞状态，底层计数器为0，接触阻塞
   fmt.Println("main .... 接触阻塞。。。。结束了。。。。")
}

func printNum1(wg  *sync.WaitGroup){
	rand.Seed(time.Now().UnixNano())
	for i:=1;i<=100;i++{
		fmt.Println("子goroutine1,i:",i)
		time.Sleep(time.Duration(rand.Intn(1000)))
	}
	wg.Done() // 计数器减一
}

func printNum2(wg  *sync.WaitGroup){
	for j:=1;j<=100;j++ {
		fmt.Println("\t子goroutine2,j:",j)
		time.Sleep(time.Duration(rand.Intn(1000)))
	}
	wg.Done() // 计数器减一
}