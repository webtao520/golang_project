/*
   通过type方式实现value.(type) 类型断言
*/

package main

import (
	"fmt"
	"strconv"
)

type Element interface{}

type List []Element

type Person struct {
	name string
	age  int
}

// 打印
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + "years)" // 将整型转换为string
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		switch value := element.(type) { // element.(type)语法不能在switch外的任何逻辑里面使用，如果你要在switch外面判断一个类型就使用comma-ok
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}
}
