package main 


import (
	"fmt"
)


// 结构体
type person struct {
	name string
	age int 
	gender string
	hobby  []string
}

func main(){
	// 声明一个person类型的变量p
	var p person
	// 通过字段赋值
	p.name="xxx"
	p.age=200
	p.gender="男"
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p) // {xxx 200 男 [篮球 足球 双色球]}
    	// 访问变量p的字段
	fmt.Printf("%T\n", p)
	fmt.Println(p.name)

	var p2 person
	p2.name = "理想"
	p2.age = 18
	fmt.Printf("type:%T value:%v\n", p2, p2)

	// 匿名结构体，多用于临时对象
	var s struct {
		x string
		y int 
	}

	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s)




}