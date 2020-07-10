package main

import (
	"fmt"
	"time"
)

func main () {
	var test=map[string]int {"a":1,"b":2}
	for _,v:=range test{
		// if  time.After(time.Second * 10) && v == 1 {
		// 	fmt.Println("----")
		// }
		select {
		case <-time.After(time.Second * 1):
			fmt.Println("----",v)
		}
	}

}