package main 

import (
	"fmt"
)

// 车轮
type Wheel  struct {
	Size int
}

//  引擎
type Engine  struct {
	 Power int    // 功率
	 Type string // 类型
}

//车
type Car struct {
	Wheel	
	Engine
}


func main(){
	/*
   c:=Car {
		 // 初始化轮子
		 Wheel:Wheel{
			 Size: 18,
		 },
		 // 初始化引擎
		 Engine:Engine{
			Type:  "1.4T",
            Power: 143,
		 }, 
   }
   fmt.Printf("%+v\n", c)
   */
   c:=new(Car)
   /*
   c.Wheel.Size=33
  c.Engine.Power=33
  c.Engine.Type="d"
  */
  c.Wheel=Wheel{Size: 34}  
  c.Engine=Engine{Power: 34,Type:"33"}
  fmt.Println(c)
}

/**
PS D:\goLang\github\golang_project\Go语言结构体\初始化内嵌结构体> go run 1.go
{Wheel:{Size:18} Engine:{Power:143 Type:1.4T}}
*/