package main 

import (
	"fmt"
)

func main (){
	a := []int{1, 2, 3}
	//copy(a, a[1:] // 2 3 3
	a = a[:copy(a, a[1:])] // 删除开头1个元素
		
	fmt.Println(a)
	//a = a[:copy(a, a[N:])] // 删除开头N个元素
}

/**
PS D:\golang\github\golang_project\go容器\Go语言从切片中删除元素> go run 3.go
[2 3]

结论：copy复制会比等号复制慢。但是copy复制为值复制，改变原切片的值不会影响新切片。
而等号复制为指针复制，改变原切片或新切片都会对另一个产生影响。
*/