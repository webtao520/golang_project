package main 

import (
	"fmt"
	"strconv"
)

func main (){
	// 字符串 转 整数
  n,err:=strconv.Atoi("lamp")
  if err !=nil {
	  fmt.Println("转换错误", err)
  }	else {
	  fmt.Println("转成的结果是", n)
  }
}
