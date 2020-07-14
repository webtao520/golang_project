package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03(){
	str:=""
	for i:=0;i<100000000;i++ {
		str+="hello" + strconv.Itoa(i)
	}
	fmt.Println(str)
}

func main (){
  //  执行test03 ,先获取到当前的unix时间戳
  start:=time.Now().Unix()
  test03()
  end:=time.Now().Unix()
  fmt.Printf("执行test03()耗费时间为%v秒\n",end-start)
}