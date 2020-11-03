// 这个示例程序展示如何用无缓冲的通道来模拟
// 2 个goroutine 间的网球比赛
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// wg 用来等待程序结束
var wg sync.WaitGroup

func init() {
    rand.Seed(time.Now().UnixNano())
}

// main 是所有Go 程序的入口
func main() {
    // 创建一个无缓冲的通道
    court := make(chan int)

    // 计数加 2，表示要等待两个goroutine
    wg.Add(2)

    // 启动两个选手
    go player("Nadal", court)
    go player("Djokovic", court)

    // 发球
    court <- 1

    // 等待游戏结束
    wg.Wait()
}

// player 模拟一个选手在打网球
func player(name string, court chan int) {
    // 在函数退出时调用Done 来通知main 函数工作已经完成
    defer wg.Done()

    for {
        // 等待球被击打过来
        ball, ok := <-court
        if !ok {
            // 如果通道被关闭，我们就赢了
            fmt.Printf("Player %s Won\n", name)
            return
        }

        // 选随机数，然后用这个数来判断我们是否丢球
        n := rand.Intn(100)
        if n%13 == 0 {
            fmt.Printf("Player %s Missed\n", name)

            // 关闭通道，表示我们输了
            close(court)
            return
        }

        // 显示击球数，并将击球数加1
        fmt.Printf("Player %s Hit %d\n", name, ball)
        ball++

        // 将球打向对手
        court <- ball
    }
}



/**
PS D:\goLang\github\golang_project\Go语言并发\Go语言无缓冲的通道> go run 1.go
Player Djokovic Hit 1
Player Nadal Hit 2
Player Djokovic Hit 3
Player Nadal Hit 4
Player Djokovic Hit 5
Player Nadal Hit 6
Player Djokovic Hit 7
Player Nadal Hit 8
Player Djokovic Missed
Player Nadal Won
*/