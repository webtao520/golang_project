package main

import (
	"fmt"
)

func main() {
	for _, r := range "hello，世界" { // 遍历的是unicode字符，不是字节，

		fmt.Printf("%T\n", r) // int32 的别名 rune


	}
	str:="lamp"
	for _, r := range str { // 遍历的是unicode字符，不是字节，

		fmt.Printf("%T\n", r) // int32 的别名 rune


	}

	//fmt.Println(utf8.RuneCountInString(str));
}

//获取字符串字符数：utf8.RuneCountInString(str) (最快)，len([]rune(str))

//拼接字符串：使用bytes.Buffer最快，Strings.Join()，+操作符
