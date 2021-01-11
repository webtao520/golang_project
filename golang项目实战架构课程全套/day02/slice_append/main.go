package main 

import (
	"fmt"
)

// append ()  为切片追加元素

func main(){
	/*
	s1 := []string{"北京", "上海", "深圳",""}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n",s1,len(s1),cap(s1)) // s1=[北京 上海 深圳] len(s1)=3 cap(s1)=3
	s1[3]="日本" // 错误的写法 会导致编译错误：索引越界
	fmt.Println(s3)
	*/
  // var s1 []string
   //fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n",s1,len(s1),cap(s1)) // s1=[北京 上海 深圳] len(s1)=3 cap(s1)=3
   //fmt.Println(s1)

   s1 := []string{"北京", "上海", "深圳"}
   fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

   // 调用append 函数必须用原来的切片遍历接受返回值
   // append 追加元素。原来的底层数组放不下的时候，go底层就会把底层数组换一个
   // 必须用变量接受append的放回值
   s1=append(s1, "日本")
   fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
   ss := []string{"武汉", "西安", "苏州"}
   //ss[3]="全椒"  错误的写法 会导致编译错误：索引越界
   s1=append(s1, ss...) // ... 表示拆开
   fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

   

   var s5 []string
   s6:=[]string{"a","b"}
   fmt.Println(append(s5,s6...))
}


/*
PS D:\golang\github\golang_project\golang项目实战架构课程全套\day02\slice_append> go run main.go
s1=[北京 上海 深圳] len(s1)=3 cap(s1)=3
s1=[北京 上海 深圳 日本] len(s1)=4 cap(s1)=6
s1=[北京 上海 深圳 日本 武汉 西安 苏州] len(s1)=7 cap(s1)=12
*/