package main

import (
	"fmt"
    "time"
)

func main() {
    stop := make(chan bool)
    // 开启goroutine
    go func() {
        for {
            select {
            case <- stop:
                fmt.Println("任务1 结束了。。。")
            default:
                fmt.Println(" 任务1 正在运行中。")
                time.Sleep(time.Second * 2)
            }
        }
    }()
    
    // 运行10s后停止
    time.Sleep(time.Second * 10)
    fmt.Println("需要停止任务1。。。")
  	stop <- true
    time.Sleep(time.Second * 1)
}

/*
		------------------执行结果---------------------------------
				任务1 正在运行中...
				任务1 正在运行中...
				任务1 正在运行中...
				任务1 正在运行中...
				任务1 正在运行中...
				任务1 正在运行中...
				需要停止任务1...
				任务1 结束了...
				任务1 正在运行中...
				任务1 正在运行中...
*/