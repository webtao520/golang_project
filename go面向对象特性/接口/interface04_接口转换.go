package main

/*
子集和超集区别???

子集定义
对于两dao个集du合A与B，如果集合A的任何一个zhi元素都是集合B的元dao素，内我们就说集合A包含于集容合B，
或集合B包含集合A，也说集合A是集合B的子集。如果集合A的任何一个元素都是集合B的元素，而集合B中至少有一个元素不属于集合A，
则称集合A是集合B的真子集。空集是任何集合的子集。 任何一个集合是它本身的子集.空集是任何非空集合的真子集.
超集定义
如果一个集合S2中的每一个元素都在集合S1中，且集合S1中可能包含S2中没有的元素，则集合S1就是S2的一个超集。
*/

import "fmt"

// 定义接口类型
type Humaner interface { // 子集
	// 方法，只有声明，没有实现，由别的类型(自定义类型)实现
	sayHi()
}

type Personer interface { // 超集
	Humaner // 匿名字段，继承了sayHi()
	sing(lrc string)
}

// 定义一个类型实现接口
type Student struct {
	name string
	id   int
}

// Student 实现了此方法
func (tmp *Student) sayHi() {
	fmt.Printf("Student[%s,%d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Student) sing(lrc string) {
	fmt.Println("Student 在唱着: ", lrc)
}

func main() {
	// 超集可以转换为子集，反过来不可以
	var iPro Personer // 超集
	var i Humaner     // 子集
	iPro = &Student{"tom", 5555}

	//iPro = i // err

	i = iPro // 可以，超集可以转换为子集
	i.sayHi()

}

/*
PS D:\goLang\github\golang_project\go面向对象特性\接口> go run .\interface04_接口转换.go
			Student[tom,5555] sayhi
*/
