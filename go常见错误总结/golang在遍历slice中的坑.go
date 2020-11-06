package main

import "fmt"

func main() {
	pase_student()
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	/*
			for range是值拷贝
		所以val这个临时变量的地址固定不变的，值在变
	*/
	for _, stu := range stus {
		m[stu.Name] = &stu // 将临时变量stu的地址给map，最后所有的地址都为同样的一个值 stu是值拷贝，而且这里是一个临时变量，用的同一个地址的空间
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

// 以上这种写法是有问题的

/**
PS D:\goLang\github\golang_project\go常见错误总结> go run golang在遍历slice中的坑.go
zhou &{wang 22}
li &{wang 22}
wang &{wang 22}
*/
