package main

import "fmt"

type Student struct {
	name string
	id   int
}

func switch_type(i []interface{}) {
	// 类型查询，类型断言
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d] 类型为int, 内容为%d\n", index, value)
		case string:
			fmt.Printf("x[%d] 类型为string, 内容为%s\n", index, value)
		case Student:
			fmt.Printf("x[%d] 类型为Student, 内容为name = %s, id = %d\n", index, value.name, value.id)
		}
	}
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1                    // int
	i[1] = "hello go"           // string
	i[2] = Student{"kim", 5555} // 结构体 Student
	switch_type(i)
}

/*
PS D:\goLang\github\golang_project\go面向对象特性\接口类型查询[断言]> go run .\switch断言.go
x[0] 类型为int, 内容为1
x[1] 类型为string, 内容为hello go
x[2] 类型为Student, 内容为name = kim, id = 5555
*/
