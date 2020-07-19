package main 

import (
	"fmt"
)

// 定义接口类型
type Humaner interface {
	// 方法，只有声明，没有实现，由别的类型(自定义类型)实现
	sayHi()
}

type Student struct {
	name string
	id int
}

type Teacher struct {
	addr string
	group string
}

type Mystr string
 
// Teacher 实现了此方法
func (tmp *Teacher) sayHi(){
	fmt.Printf("Teacher[%s,%s] sayhi\n",tmp.addr,tmp.group)
}

// Student 实现了此方法
func (tmp *Student)sayHi(){
	fmt.Printf("Student[%s,%d] sayhi\n",tmp.name,tmp.id)
}

// Mystr 实现此方法
func (tmp *Mystr) sayHi(){
	fmt.Printf("Mystr[%s] sayHi\n",*tmp)
}

func main () {
   // 定义接口类型的变量
   var i Humaner
   // 只要实现了此接口方法的类型，那么这个类型的变量(接受者类型)就可以给i赋值
   s:=&Student{
	   "tom",
	   555,
   }
   i=s 
   i.sayHi()

   t:=&Teacher{
	  "bj",
	  "go",
   }
   i=t
   i.sayHi()

   var str Mystr = "hello golang"
   i=&str
   i.sayHi()
}

/*
	PS D:\goLang\github\golang_project\go面向对象特性> go run .\interface01.go
			Student[tom,555] sayhi
			Teacher[bj,go] sayhi
			Mystr[hello golang] sayHi
*/