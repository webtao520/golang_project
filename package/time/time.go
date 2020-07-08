package main 


import (
	"fmt"
	"time"
)

func main (){
	// 获取当前时间
	now:=time.Now()
	fmt.Printf("now=%v now type=%T\n",now,now)

	// 通过now可以获取到年月日，时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n",now.Month())
	fmt.Printf("月=%v\n",int(now.Month()))
	fmt.Printf("日=%v\n",now.Day())
	fmt.Printf("时=%v\n",now.Hour())
	fmt.Printf("分=%v\n",now.Minute())
	fmt.Printf("秒=%v\n",now.Second())

}