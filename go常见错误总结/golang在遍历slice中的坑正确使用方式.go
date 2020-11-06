// 应该修改下面的方式， 确保在遍历的时候不会出现地址都是同一个值
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
	for i, v := range stus {
		fmt.Println("key===", i)

		m[v.Name] = &stus[i]
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
