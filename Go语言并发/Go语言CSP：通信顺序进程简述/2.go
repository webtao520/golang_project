package  main 

import (
	"fmt"
	"sync"
)

func main(){
	var mu sync.Mutex 

	mu.Lock()
	go func (){
		fmt.Println("C语言中文网")
		mu.Unlock()
	}()

	mu.Lock()
}