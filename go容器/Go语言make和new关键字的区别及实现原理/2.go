package main

import (
	"fmt"
)

type Student struct {
	name string
	age int
 }
 
 func main (){
	var s *Student
	s = new(Student)   //分配空间  类型的指针 指针指向分配内存的地址
	s.name ="dequan"
	fmt.Println(s)
 }



 /**
	PS D:\goLang\github\golang_project\go容器\Go语言make和new关键字的区别及实现原理> go run 2.go
	&{dequan 0}
 */

