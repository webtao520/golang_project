package main 


import (
	"fmt"
	"time"
)

// work pool 
// `goroutine`池 <-chan int 单向通道， 只读/接收  chan<-int 只写/ 发送
func worker(id int, jobs <-chan int, results chan<-int){
   for j:=range jobs {
	fmt.Printf("worker:%d start job:%d\n", id, j)
	time.Sleep(time.Second)
	fmt.Printf("worker:%d end job:%d\n", id, j)
	results <- j * 2
   }
}

func main (){
	 // 初始化带缓冲区的通道
	 jobs := make(chan int, 100)
	 results := make(chan int, 100)
	 // 开启3个goroutine
	 for w := 1; w <= 3; w++ {
		 go worker(w, jobs, results)
	 }
	 // 5个任务
	 for j := 1; j <= 5; j++ {
		 jobs <- j
	 }
	 close(jobs)
	 // 输出结果
	 for a := 1; a <= 5; a++ {
		 <-results
	 }

}