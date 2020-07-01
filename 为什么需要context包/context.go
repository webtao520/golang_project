package main

import (
	  "fmt"
    "sync"
    "time"
    "strconv"
)

var wg sync.WaitGroup

func run(task string) {
    fmt.Println(task, "start。。。")
    time.Sleep(time.Second * 2)
    // 每个goroutine运行完毕后就释放等待组的计数器
    wg.Done()
}

func main() {
    wg.Add(2)			// 需要开启几个goroutine就给等待组的计数器赋值为多少，这里为2
    for i := 1; i < 3; i++ {
        taskName := "task" + strconv.Itoa(i)
				go run(taskName)
    }
    // 等待，等待所有的任务都释放 等待组计数器值为 0
    wg.Wait()
    fmt.Println("所有任务结束。。。")
}

/*
  -----------------------运行结果----------------------------
				task2 start。。。
				task1 start。。。
				所有任务结束。。。
*/