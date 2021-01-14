package main 


import (
	"fmt"
)

// 接口的实现
type animal interface {
	 move()
	 eat(string)
}

type cat struct {
	name string
	feet int8 
}

func(c cat) move(){
	 fmt.Println("走猫步....")
}

func (c cat) eat(food string){
	fmt.Printf("猫吃%s...\n", food)
}


type chicken  struct{
	feet int8 
}

func (c chicken) move() {
	fmt.Println("鸡动!")
}

func (c chicken) eat(food string) {
	fmt.Printf("吃%s...\n", food)
}


func main(){
   var a1 animal //定义一个接口类型的变量
   fmt.Printf("%T\n", a1) // <nil>

   bc := cat{ // 定义一个cat类型的变量bc
	name: "淘气",
	feet: 4,
}
   a1 = bc
   a1.eat("小黄鱼")
   fmt.Printf("%T\n",a1)
   fmt.Println(a1)
 
   kfc := chicken{
	feet: 2,
	}
	a1 = kfc
	fmt.Printf("%T\n", a1)
}

/*
PS D:\goLang\github\golang_project\golang项目实战架构\day05\interface3> go run main.go
<nil>
猫吃小黄鱼...
main.cat
{淘气 4}
main.chicken
*/