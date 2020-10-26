package main 

import (
	"fmt"
)

func main(){
	a:=[]int{1,2,3}
	a=a[1:]    // 删除开头1个元素
	//a=a[N:] // 删除开头N个元素
	fmt.Println(a)
}

/*
	PS D:\golang\github\golang_project\go容器\Go语言从切片中删除元素> go run 1.go
	[2 3]
*/