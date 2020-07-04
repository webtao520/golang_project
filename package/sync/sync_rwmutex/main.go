package main 

/*
		并发读写map
*/

import (
	"fmt"
	"sync"
	"time"
)

func main (){
	var rwm sync.RWMutex
		map1:=make(map[string]string)
		for i:=1;i<=100;i++ {
			go func (i int) {
				rwm.Lock()
				map1[fmt.Sprintf("key,%d",i)] = fmt.Sprintf("数据%d",i)
				rwm.Unlock()
			}(i)
		}
		time.Sleep(2*time.Second)
		fmt.Println(map1)
}