package main 

import (
	"fmt"
)

// map 和 slice组合

func main(){
	// 元素类型为map的切片 
	var s1 = make([]map[int]string,10,10)
	// 没有对内部的map做初始化
	s1[0] = make(map[int]string,1)
	s1[0][10] = "少喝"
	// [map[10:少喝] map[] map[] map[] map[] map[] map[] map[] map[] map[]]
	fmt.Println(s1)
	// 值为切片类型的map
	var m1 = make (map[string][]int,10)
	m1["北京"] = []int{1,2,3}
	fmt.Println(m1) // map[北京:[1 2 3]]

}