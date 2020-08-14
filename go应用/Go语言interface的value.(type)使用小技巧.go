/*
	interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。
	空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。
	我们知道interface的变量里面可以存储任意类型的数值(该类型实现了interface)。
	那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？目前常用的有两种方法：
	Comma-ok断言

	Go语言里面有一个语法，可以直接判断是否是该类型的变量：

	value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。

	如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。

	让我们通过一个例子来更加深入的理解
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

//定义了String方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}
func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "hello"
	list[2] = Person{
		name: "zhangtao",
		age:  40,
	}
	for index, element := range list {
		/*
			Go语言里面有一个语法，可以直接判断是否是该类型的变量：
			value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
		*/
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Println("list[%d] is a Person and its value is %s\n", index, value)
		}

	}
}

/*
	list[0] is an int and its value is 1
	list[1] is a string and its value is hello
	list[2] is a Person and its value is (name: zhangtao - age: 40 years)
*/
