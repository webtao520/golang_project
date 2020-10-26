package main 

import (
	"fmt"
)

func main(){
	a:=[]int{1,2,3}
	a=a[:len(a)-1] // 删除尾部1个元素
	fmt.Println(a)
}

/**
PS D:\goLang\github\golang_project\go容器\Go语言从切片中删除元素> go run 5.go
[1 2]


a = []int{1, 2, 3}
a = a[:len(a)-1] // 删除尾部1个元素
a = a[:len(a)-N] // 删除尾部N个元素
*/