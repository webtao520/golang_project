package main 

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
sync包提供了基本的同步基元，如互斥锁。除了Once和WaitGroup类型，
大部分都是适用于低水平程序线程，
高水平的同步使用channel通信更好一些。

本包的类型的值不应被拷贝。
*/

// waitGroup
func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()    // int64
		r2 := rand.Intn(10) // 0<= x < 10
		fmt.Println(0-r1, 0-r2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	// f()
	// wg.Add(10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	// ?如何知道这10个goroutine都结束了
	// time.Sleep(?)
	wg.Wait() // 等待wg的计数器减为0
}
