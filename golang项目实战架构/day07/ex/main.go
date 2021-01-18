package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用goroutine和channel实现一个计算int64随机数各位数和的程序。
1. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
2. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
3. 主goroutine从resultChan取出结果并打印到终端输出
*/

// job ...
type job struct {
	 value int64
}

// result..
type result struct {
	 job *job 
	 sum  int64
}


//  初始化chan  带缓冲区的通道
var jobChan  =make(chan *job,100)
var resultChan = make(chan *result,100)
// goroutine 计数器
var wg sync.WaitGroup


// 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
// 单向通道  保证参数在函数间传递的安全性
func zhoulin(zl  chan<- *job){
	defer wg.Done()
	// 循环生成int64类型的随机数，发送到jobChan
	for {
		 x:=rand.Int63()
		 newJob:=&job{
			value:x, 
		 }
		 zl <-newJob
		 time.Sleep(time.Microsecond*500)
	}
}

// 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
func baodelu(zl <-chan *job, resultChan chan<- *result){
	defer wg.Done()
	// 从jobchan 中取出随机数计算各位数的和，将结果放到resultChan
	for {
		 job :=<-zl
		 sum:=int64(0) // 初始化sum = 0
		 n:=job.value
		 for n>0 {
			 sum+=n % 10
			 n=n / 10
		 }
		 newResult:=&result{
			 job:job,
			 sum:sum,
		 }
		 resultChan <- newResult
	}
}

func main(){
	wg.Add(1)
	go zhoulin(jobChan)
	// 开启24 个goroutine执行baodelu
	for i:=0;i<24;i++{
		 go baodelu(jobChan,resultChan)
	}
	// 主goroutine从resultChan取出结果并打印到终端输出
	for result:=range resultChan{
		fmt.Printf("value:%d sum:%d\n",result.job.value,result.sum)
	}
	wg.Wait()
}