package main

import (
	"chapter08/clsfactory/base"
	_ "chapter08/clsfactory/cls1" // 匿名引用cls1包，自动注册
	_ "chapter08/clsfactory/cls2" // 匿名引用cls2包，自动注册
)

func main() {

	// 根据字符串动态创建一个Class1实例
	c1 := base.Create("Class1")
	c1.Do()

	// 根据字符串动态创建一个Class2实例
	c2 := base.Create("Class2")
	c2.Do()

}
