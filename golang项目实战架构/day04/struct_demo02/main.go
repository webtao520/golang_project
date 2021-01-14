package main 

import (
	"fmt"
)

// 结构体占用一块连续的内存空间
/*
		Byte是字节，是计算机文件的最小单位； 
		bit是位，2进制运算符，一个字节占8位  bit是字bai位，也就是计算du机存储数据的“0”、“1”
*/
type x struct {
	a int8 // 8bit => 1byte
	b int8
	c int8
}

func  main (){
   m :=  x{
		a:int8(8),
		b:int8(20),
		c:int8(30),
   }
   fmt.Printf("%p\n", &(m.a))
   fmt.Printf("%p\n", &(m.b))
   fmt.Printf("%p\n", &(m.c))
}

/**
0xc00000a0b0
0xc00000a0b1
0xc00000a0b2
*/