package  main 

import (
	"fmt"
)


func main (){
	str:="hello 世界"
	for _,r:=range str {   // 使用range，其实是使用rune类型来编码的，rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
		fmt.Printf("%c",r)
	}
}