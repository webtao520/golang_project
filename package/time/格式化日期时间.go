package main

import (
	"fmt"
	"time"
)

func main (){
	now:=time.Now()
	fmt.Printf("当前年月日 %d - %d - %d %d : %d : %d \n", now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	
	dataStr:=fmt.Sprintf("当前年月日 %d - %d -%d %d : %d : %d",now.Year(), now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Printf("dataStr=%v\n", dataStr)
}