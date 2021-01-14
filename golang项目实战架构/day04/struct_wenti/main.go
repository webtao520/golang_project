package main 


import (
	"fmt"
)

// 结构体遇到的问题
// 1. myInt(100) 是个啥？
type myInt int 

func (m myInt) hello(){
	fmt.Println("我是一个自定义类型int")
}

type person struct {
	name string
	age int 
}

func main(){
	// 声明一个int32类型的变量x,它的值是10
	// 方法1:
	// var x int32
	// x = 10
	// 方法2:
	// var x int32 = 10
	// 方法3:
	// var x = int32(10)
	// 方法4:
	// x := int32(10)
	// fmt.Println(x)

		// 声明一个myInt类型的变量m,它的值是100
	// 方法1:
	// var m myInt
	// m = 100
	// fmt.Println(m)
	// 方法2：
	//  var m myInt = 100
	 // 方法3：
	//  var m = myInt(100)
	 // 方法4：
	//  m:=myInt(100) // 强制类型转换
	//  fmt.Println(m)
	//  m.hello()


	// 问题2： 结构体初始化
	// 方法1：
	var p person // 声明一个person类型的变量p
	p.name="元素"
	p.age=19
	fmt.Println(p)
	var p1 person
	p1.name = "周林"
	p1.age = 9000
	fmt.Println(p1)

	// 方法2：
	s1:=[]int{1,2,3,4}
	m1:=map[string]int{
		"stu":100,
		"stu2":99,
		"stu3":0,
	}
	fmt.Println(s1, m1)
	// 键值初始化
	//var p2 person
	var p2 = person{
		name: "冠华",
		age:  15,
	}
	fmt.Println(p2)

	// 值列表初始化
	var p3 = person{
		"理想",
		100,
	}
	fmt.Println(p3)
	
}


// 问题3:为什么要有构造函数
func newPerson(name string, age int) person {
	// 别人调用我,我能给她一个person类型的变量
	return person{
		name: name,
		age:  age,
	}
}

// func newPerson(name string, age int) *person {
// 	// 别人调用我,我能给她一个person类型的变量
// 	return &person{
// 		name: name,
// 		age:  age,
// 	}
// }
