package main 

import (
	"fmt"
)

 // 单行注释

/*
多行注释
*/


 // Go语言函数外的语句必须以关键字开头
var name = "娜扎"
var age int

const (
	Num = 100
)


// main函数是入口函数
// 它没有参数也没有返回值
func main(){
	// 函数内部定义的变量必须使用
     var isOK = true 
	 fmt.Println(isOK)

	  // for 1 
	  for i:=0;i<10;i++{
		  fmt.Println(i)
	  }

	  // for 2
	 var i = 0
	 for ;i<10;i++{
		 fmt.Println(i)
	 }

	 // for 3
	 var j = 0
	 for j<10{
		 fmt.Println(j)
		 j++
	 }
	 // for 4
	// for {
	// 	fmt.Println("无限循环")
	// }
	//for 5
	s := "hello"
	for i,v:=range s{
		fmt.Println(i, v)
		fmt.Printf("%d %c\n", i, v) //  相应Unicode码点所表示的字符 
	}

	var f1 = 1.234 // float64
	fmt.Printf("%T\n", f1)
	var f2 = float32(f1)
	fmt.Printf("%T\n", f2)
	//fmt.Println(f1 == f2)


		// byte 和 rune 非ascii吗
	s1 := "Hello"
	s2 := "沙河有沙又有河"
	for _, v := range s1 { 
		fmt.Printf("%c %T\n", v, v)    //  %c 相应 Unicode 码点所表示的字符 
		fmt.Println("===>",string(v))
	}

	for _, v := range s2 {
		fmt.Printf("%c %T\n", v, v)   //  %c 相应 Unicode 码点所表示的字符  
		fmt.Println("===>>",v)
	}

		// 八进制数
		var n1 = 0777
		// 十六进制数
		var n2 = 0xff
		fmt.Println(n1, n2)
		fmt.Printf("%o\n", n1) //  777
		fmt.Printf("%x\n", n2) // ff
	

}