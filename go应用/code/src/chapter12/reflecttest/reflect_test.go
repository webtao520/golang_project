package main

import (
	"reflect"
	"testing"
)

// 声明一个结构体，拥有1个字段
type data struct {
	Hp int
}

func BenchmarkNativeAssign(b *testing.B) {

	// 实例化结构体
	v := data{Hp: 2}

	// 停止基准测试的计时器
	b.StopTimer()
	// 重置基准测试计时器数据
	b.ResetTimer()

	// 重新启动基准测试计时器
	b.StartTimer()

	// 根据基准测试数据进行循环测试
	for i := 0; i < b.N; i++ {

		// 结构体成员赋值测试
		v.Hp = 3
	}

}

func BenchmarkReflectAssign(b *testing.B) {

	v := data{Hp: 2}

	// 取出结构体指针的反射值对象，并取其元素
	vv := reflect.ValueOf(&v).Elem()

	// 根据名字取结构体成员
	f := vv.FieldByName("Hp")

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {

		// 反射测试设置成员值性能
		f.SetInt(3)
	}
}

func BenchmarkReflectFindFieldAndAssign(b *testing.B) {

	v := data{Hp: 2}

	vv := reflect.ValueOf(&v).Elem()

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {

		// 测试结构体成员的查找和设置成员的性能
		vv.FieldByName("Hp").SetInt(3)
	}

}

func foo(v int) {

}

func BenchmarkNativeCall(b *testing.B) {

	for i := 0; i < b.N; i++ {

		// 原生函数调用
		foo(0)
	}
}

func BenchmarkReflectCall(b *testing.B) {

	// 取函数的反射值对象
	v := reflect.ValueOf(foo)

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		// 反射调用函数
		v.Call([]reflect.Value{reflect.ValueOf(2)})
	}

}
