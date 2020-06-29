package main 

import (
	"fmt"
	"sync"
)


/*
    在解锁之前加锁会导致死锁
*/
func main (){
	var mutex sync.Mutex
	mutex.Lock()
	fmt.Println("Locked")
	mutex.Lock() // Locked fatal error: all goroutines are asleep - deadlock!
}