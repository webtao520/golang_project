package main 

import (
	"fmt"
)

func yuanshuai(name string) {
	fmt.Println("Hello",name)
}

//函数作为参数
func lixiang(f func(string), name string){
	f(name)
}

//函数作为返回值
func zhoulin() func(int,int) int{
  return func(x,y int) int{
    return x+y
  }
}

func low(f func()){
	f()
}

// 闭包 = 函数 + 外部变量
func bi(f func(string),name string) func(){
    return func() {
		f(name)
	}
}

func main(){
	 lixiang(yuanshuai,"理想")

	 ret:=zhoulin()
	 fmt.Printf("%T\n", ret) // func(int, int) int
	 sum := ret(10, 20)
	 fmt.Println(sum) // 30
	 	// 闭包示例
	fc := bi(yuanshuai, "元帅")
	low(fc)
}