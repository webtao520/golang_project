//下面使用结构体定义一个命令行指令（Command），指令中包含名称、变量关联和注释等，对 Command 进行指针地址的实例化，并完成赋值过程，代码如下
package  main 

import (
	"fmt"
)

type Command  struct {
	Name string  // 指令名称
	Var *int // 指令绑定的变量
	Comment string    // 指令的注释
}

var version  int =1 

func main(){
   cmd:=&Command{} // 对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作
   cmd.Name="version"
   cmd.Var= &version
   cmd.Comment="show version"
   fmt.Println(cmd)
}

/**
PS D:\goLang\github\golang_project\Go语言结构体\Go语言实例化结构体> go run  2.go
&{version 0x563230 show version}
*/