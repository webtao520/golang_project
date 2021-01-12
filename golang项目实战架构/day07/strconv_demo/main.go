package main 

import (
	"fmt"
	"strconv" // strconv包实现了基本数据类型和其字符串表示的相互转换。
)

// strconv

func main(){
   // 从字符串中解析出对应的数据
   str:="100000"
   //ret1:=int64(str)
   ret1,err:=strconv.ParseInt(str,10,64) // 返回字符串表示的整数值，接受正负号。
	if err !=nil {
		fmt.Println("parseint failed, err:", err)
		return
	}
	fmt.Printf("%#v %T\n",ret1,int(ret1))
	

	// 将字符串 转换成整型 Atoi  Atoi是ParseInt(s, 10, 0)的简写。
	retInt,_:=strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)

	// 从字符串中解析出布尔值
	boolStr := "true"
	// 返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
	boolValue ,_:=strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n",boolValue,boolValue) // true bool

	//  把数组转换成字符串类型
	i:=98
	//ret2:=string(i)
	// ret2 := string(i) // "a"
	ret2:=fmt.Sprintf("%d",i) 
   //fmt.Printf("%T\n",ret2) // string
   fmt.Printf("%#v", ret2) // "98"

   ret3 := strconv.Itoa(i) // Itoa是FormatInt(i, 10) 的简写。
   fmt.Printf("%#v", ret3)

}