package main 

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main (){
  var wg sync.WaitGroup  // 同步等待组
  fmt.Printf("%T\n",wg)
  wg.Add(2) // 设置计数器的数量	
  go p1(&wg)
  go p2(&wg)
  wg.Wait() //进入阻塞状态，当计数器为0，接触阻塞
  fmt.Println("main.....接触阻塞...结束了") 
}

func p1(wg *sync.WaitGroup){
  rand.Seed(time.Now().UnixNano())
  for i:=0;i<=100;i++ {
	  fmt.Println("子goroutine1,i:", i)
	  time.Sleep(time.Duration(rand.Intn(1000)))
  }
  wg.Done() // 将计数器减1 Add(-1)
}

func p2(wg *sync.WaitGroup){
  for j:=1;j<=100;j++ {
	  fmt.Println("\t子goroutine2,j:",j)
	  time.Sleep(time.Duration(rand.Intn(1000)))
  }
  wg.Done()
}