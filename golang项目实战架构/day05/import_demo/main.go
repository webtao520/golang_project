package main

import (
	"fmt"
	zhoulin	"../calc"
)

var x =100
const pi = 3.12

func init(){
	fmt.Println("自动执行!")
	fmt.Println(x, pi)
}

func main(){
	ret := zhoulin.Add(10, 20)
	fmt.Println(ret)  
}

/**
import 我时自动执行...
自动执行!
100 3.12
30
*/