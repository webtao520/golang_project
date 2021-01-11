package main

import (
	"fmt"
)

// byte uint8  和 rune int32 类型
// go 语言中为了处理非ascii编码类型的字符，定义新的rune 类型
func main() {
	s := "Hello沙河사샤"
	// len() 求得是 byte 字节数量
	n := len(s)    // 求字符串s的长度，把长度保存到变量n中
	fmt.Println(n) // 17

	/*
		for i := 0; i < len(s); i++ {
			//fmt.Println(s[i])
			fmt.Printf("%c\n", s[i]) // 相应 Unicode码点所表示的字符   %c:字符
		}
	*/
	/*
		for _, c := range s { // 从字符串拿出来具体的字符
			fmt.Printf("%c\n", c) //  %c:字符
		}
	*/

	// 字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) // 把字符串强制转换成了一rune切片
	s3[0] = '红'
	fmt.Println(string(s3)) // 把 rune  切片强制转换成字符串

	c1 := "红"
	c2 := '红'                           // rune(int32)
	fmt.Printf("c1:%T c2:%T\n", c1, c2) // c1:string c2:int32 (rune)
	c3 := "H"                           // string
	c4 := byte('H')                     // byte (uint8)
	fmt.Printf("c3:%T c4:%T\n", c3, c4) // c3:string c4:uint8

	// 类型转换
	n1 := 10 // int
	var f float64
	f = float64(n1)
	fmt.Println(f)        //  10
	fmt.Printf("%T\n", f) // float64

}
