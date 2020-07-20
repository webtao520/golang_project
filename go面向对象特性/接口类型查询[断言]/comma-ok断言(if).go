package main 

import (
	"fmt"
)

type Student struct {
	 name string 
	 id int 
}

func main (){
	i:=make([]interface{},3)
	i[0]=1 // int
	i[1]="hello go" // string
	i[2]=Student{"kim",5555}
	// 类型查询，类型断言
	// 第一个返回下标，第二个返回下标对应的值 data分别是i[0],i[1],i[2]
	for index,data:=range i {
		// 第一个返回的是值，第二个返回判断你结果的真假
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为int，内容为%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为string，内容为%s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d] 类型为Student，内容为name = %s, id = %d\n", index, value.name, value.id)
		}
	}
}

/*
PS D:\goLang\github\golang_project\go面向对象特性\接口类型查询[断言]> go run '.\comma-ok断言(if).go'
		x[0] 类型为int，内容为1
		x[1] 类型为string，内容为hello go
		x[2] 类型为Student，内容为name = kim, id = 5555
*/