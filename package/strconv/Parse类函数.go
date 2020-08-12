/*
Parse类函数用于转换字符串为给定类型的值：
ParseBool()、ParseFloat()、ParseInt()、ParseUint()。

由于字符串转换为其它类型可能会失败，所以这些函数都有两个返回值，
第一个返回值保存转换后的值，第二个返回值判断是否转换成功。
*/

package main

import (
	"fmt"
	"strconv"
)

func main (){
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("-42", 10, 64)
	u, err := strconv.ParseUint("42", 10, 64)
	if err !=nil {
	    fmt.Println(err)
	}else {
		fmt.Printf("%T",b) // bool
	}
}

/*
ParseFloat()只能接收float64类型的浮点数。

ParseInt()和ParseUint()有3个参数：

func ParseInt(s string, base int, bitSize int) (i int64, err error)
func ParseUint(s string, base int, bitSize int) (uint64, error)
bitSize参数表示转换为什么位的int/uint，有效值为0、8、16、32、64。当bitSize=0的时候，表示转换为int或uint类型。例如bitSize=8表示转换后的值的类型为int8或uint8。

base参数表示以什么进制的方式去解析给定的字符串，有效值为0、2-36。当base=0的时候，表示根据string的前缀来判断以什么进制去解析：0x开头的以16进制的方式去解析，0开头的以8进制方式去解析，其它的以10进制方式解析。

以10进制方式解析"-42"，保存为int64类型：

i, _ := strconv.ParseInt("-42", 10, 64)
以5进制方式解析"23"，保存为int64类型：

i, _ := strconv.ParseInt("23", 5, 64)
println(i)    // 13
因为5进制的时候，23表示进位了2次，再加3，所以对应的十进制数为5*2+3=13。

以16进制解析23，保存为int64类型：


i, _ := strconv.ParseInt("23", 16, 64)
println(i)    // 35
因为16进制的时候，23表示进位了2次，再加3，所以对应的十进制数为16*2+3=35。

以15进制解析23，保存为int64类型：


i, _ := strconv.ParseInt("23", 15, 64)
println(i)    // 33
因为15进制的时候，23表示进位了2次，再加3，所以对应的十进制数为15*2+3=33。
*/