package main

import (
	"fmt"
)

func main (){
	var num int =9 
	fmt.Printf("num 地址 %v\n", &num)
	var ptr *int
	ptr=&num 
	*ptr=10
	fmt.Println("num = ",num)
}