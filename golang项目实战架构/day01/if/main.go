package main 


import (
	"fmt"
)

// if 条件判断

func main(){ 
	// age := 19

	// if age > 18 { // 如果 age > 18 就执行这个{}中的代码
	// 	fmt.Println("澳门首家线上赌场开业啦！")
	// } else { // 否则就执行这个{}中的代码
	// 	fmt.Println("改写暑假作业啦！")
	// }

	// 多个判断条件
	// if age > 35 {
	// 	fmt.Println("人到中年")
	// } else if age > 18 {
	// 	fmt.Println("青年")
	// } else {
	// 	fmt.Println("好好学习！")
	// }


	// 作用域
	// age 变量此时只在if条件判断语句中生效
	if age :=19;age > 18{
		fmt.Println("澳门首家线上赌场开业啦！")
	}else {
		fmt.Println("改写暑假作业啦！")
	}

	// fmt.Println(age) 在这里是找不到age

}