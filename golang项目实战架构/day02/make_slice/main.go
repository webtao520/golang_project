package main 

import (
	"fmt"
)

// make () 函数创建切片

func main (){
   s1:=make([]int,5,10)
   //s1=[]int{1,2,3,4,5,7}
   fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n",s1,len(s1),cap(s1)) // s1=[0 0 0 0 0] len(s1)=5 cap(s1)=10

   for i:=0;i<len(s1);i++{
	fmt.Println(i)
   }

   s2:=make([]int,0,10)
   fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s2, len(s2), cap(s2)) // s1=[] len(s1)=0 cap(s1)=10

   // 切片的赋值
   s3:=[]int{1,3,5}
   s4:=s3 // s3 和s4都指向了同一个底层数组
   fmt.Println(s3,s4) // [1 3 5] [1 3 5]
   s3[0]=1000  // 切片是引用类型
   fmt.Println(s3,s4)  // [1000 3 5] [1000 3 5]

   // 切片的遍历
   // 1. 索引遍历
   for i:=0;i<len(s3);i++ {
	   fmt.Println(s3[i])
   }
   //2. for range 循环
   for i,v:=range s3{
       fmt.Println(i,v)
   }
}


/**
PS D:\golang\github\golang_project\golang项目实战架构课程全套\day02\make_slice> go run  main.go
s1=[0 0 0 0 0] len(s1)=5 cap(s1)=10
0
1
2
3
4
s1=[] len(s1)=0 cap(s1)=10
[1 3 5] [1 3 5]
[1000 3 5] [1000 3 5]
1000
3
5
0 1000
1 3
2 5
*/