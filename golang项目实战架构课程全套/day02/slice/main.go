package main 


import (
	"fmt"
)

//  切片 slice

func main(){
	// 切片的定义
	var s1 []int // 定义一个存放int类型元素的切片
	var s2 []string //定义一个存放string 类型元素的切片
	fmt.Println(s1,s2) // [] []
	
	fmt.Println(s1 == nil) // true
	fmt.Println(s2 == nil)  // true 

	//  初始化 
	s1=[]int{1,2,3}
	s2 = []string{"沙河", "张江", "平山村"}
	fmt.Println(s1, s2) // [1 2 3] [沙河 张江 平山村]
	fmt.Println(s1 == nil) // false
	fmt.Println(s2 == nil) // false
	
	// 长度 和 容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 2.  由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // 基于一个数组切割，左包含右不包含，（左闭右开）
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	s5 := a1[:4] // => [0:4] [1 3 5 7]
	s6 := a1[3:] // => [3:len(a1)]  [7 9 11 13]
	s7 := a1[:]  // => [0:len(a1)]
	fmt.Println(s5, s6, s7)
		// 切片的容量是指底层数组的容量
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
		// 底层数组从切片的第一个元素到最后的元素数量
		fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))
		s8 := s6[3:] // [13]
		fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))
			// 切片是引用类型，都指向了底层的一个数组。
	fmt.Println("s6：", s6)
	a1[6] = 1300 // 修改了底层数组的值
	fmt.Println("s6：", s6)
	fmt.Println("s8：", s8)
	
}	

/*
PS D:\golang\github\golang_project\golang项目实战架构课程全套\day02\slice> go run .\main.go
[] []
true
true
[1 2 3] [沙河 张江 平山村]
false
false
len(s1):3 cap(s1):3
len(s2):3 cap(s2):3
[1 3 5 7]
[3 5 7 9 11]
[1 3 5 7] [7 9 11 13] [1 3 5 7 9 11 13]
len(s5):4 cap(s5):7
len(s6):4 cap(s6):4
len(s8):1 cap(s8):1
s6： [7 9 11 13]
s6： [7 9 11 1300]
s8： [1300]
*/