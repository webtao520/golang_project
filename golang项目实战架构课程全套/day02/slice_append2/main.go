package main 


import (
	"fmt"
)

// 关于append 删除切片中某个元素

func main(){
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	s1:=a1[:] // 切片中包含了数组中所有元素
	

	// 删除索引为1的那个3  1   5, 7, 9, 11, 13, 15, 17
	s1=append(s1[0:1],s1[2:]... )

	fmt.Println(s1)
	fmt.Println(a1) // 

}

/*
PS D:\golang\github\golang_project\golang项目实战架构课程全套\day02\slice_append2> go run .\main.go
[1 5 7 9 11 13 15 17]
[1 5 7 9 11 13 15 17 17]
*/