package main 

import (
	"fmt"
	"time"
)

func main () {
	// unix 和 unixNano 使用
	now:=time.Now()
	fmt.Printf("unix时间戳=%v unixnano 时间戳=%v\n",now.Unix(), now.UnixNano())
}