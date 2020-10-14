package main 

import (
	"fmt"
)

func main(){
	var numbers []int 
	for i:=0;i<10;i++ {
		numbers=append(numbers,i)	
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
}

/**
占位符         说明                      举例                             输出
%p      十六进制表示，前缀 0x          Printf("%p", &people)             0x4f57f0

PS D:\golang\github\golang_project\go容器\使用append()为切片添加元素> go run  2.go
len: 1  cap: 1 pointer: 0xc00000a0b0
len: 2  cap: 2 pointer: 0xc00000a0f0
len: 3  cap: 4 pointer: 0xc0000103e0
len: 4  cap: 4 pointer: 0xc0000103e0
len: 5  cap: 8 pointer: 0xc00000e280
len: 6  cap: 8 pointer: 0xc00000e280
len: 7  cap: 8 pointer: 0xc00000e280
len: 8  cap: 8 pointer: 0xc00000e280
len: 9  cap: 16 pointer: 0xc000102080
len: 10  cap: 16 pointer: 0xc000102080
*/