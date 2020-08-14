package main

import (
	"fmt"
	"time"
)

func main() {
	// unix 和 unixNano 使用
	now := time.Now()
	fmt.Printf("unix时间戳=%v unixnano 时间戳=%v\n", now.Unix(), now.UnixNano())
	LastTime := time.Now().Local().Unix()
	fmt.Println("Unix 时间戳=%v", LastTime) // Local返回采用本地和本地时区，但指向同一时间点的Time。 Unix时间
}

/*
 	############### 运行结果 ################
	unix时间戳=1597371937 unixnano 时间戳=1597371937794620000
	Unix 时间戳=%v 1597371937
*/
