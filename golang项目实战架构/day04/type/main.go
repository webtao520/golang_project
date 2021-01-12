package main 

import (
	"fmt"
	"unicode/utf8"
)

//  自定义类型和类型别名
// type 后面跟的是类型
type myInt int 
type yourInt = int  // 类型别名
/*
func main(){
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T %d\n",n,n) // main.myInt 100

	var m yourInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var c rune
	c = '中' // 字符：单引号包裹的是字符，单个字母、单个符号、单个文字
	fmt.Println(c)
	fmt.Printf("%T\n", c) // int32 rune
}


byte 等同于int8，常用来处理ascii字符
rune 等同于int32,常用来处理unicode或utf-8字符
*/

/*
func main() {

    var str = "hello 你好"
    fmt.Println("len(str):", len(str)) // len(str): 12

 // 理解 unicode 和 utf-8 编码
   golang中string底层是通过byte数组实现的。中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。

}

*/

func main() {

	var str = "hello 你好"


    //golang中string底层是通过byte数组实现的，座椅直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
    fmt.Println("len(str):", len(str))
    
    //以下两种都可以得到str的字符串长度
    
    //golang中的unicode/utf8包提供了用utf-8获取长度的方法
    fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

    //通过rune类型处理unicode字符
	fmt.Println("rune:", len([]rune(str)))
	
	for k,v:=range str {
		fmt.Println(k,"===>",v)
	}
}

/*
len(str): 12
RuneCountInString: 8
rune: 8
*/