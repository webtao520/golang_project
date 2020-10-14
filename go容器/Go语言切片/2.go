package main

import (
	"fmt"
)

func main (){
	var highRiseBuilding  [30]int 
	for i:=0;i<30;i++ {
		highRiseBuilding[i]=i+1	
	}
	// 区间
	fmt.Println(highRiseBuilding[10:15]) // [11 12 13 14 15]
	//fmt.Println(len(highRiseBuilding))

	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])

	// 开头到中间指定位置的所有元素
	fmt.Println(highRiseBuilding[:2])

	// 从开头到尾部所有元素
	fmt.Println(highRiseBuilding[1:len(highRiseBuilding)])
}

/**
PS D:\golang\github\golang_project\go容器\Go语言切片> go run 2.go
[11 12 13 14 15]
[21 22 23 24 25 26 27 28 29 30]
[1 2]
[2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30]
*/