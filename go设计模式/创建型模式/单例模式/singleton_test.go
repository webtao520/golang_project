package singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	sin1 := GetSingleton()
	sin2 := GetSingleton()

	if sin1 != sin2 {
		t.Error("实例对象不一样")
	} else {
		fmt.Println("实例对象相等")
	}
}

/*
	 ---------------------获取执行结果---------------------------
		PS D:\goLang\github\golang_project\go设计模式\单例模式> go test  -v singleton_test.go  singleton.go
				=== RUN   TestSingleton
				实例对象的信息和地址 &{12} 0x656c60
				实例对象的信息和地址 &{12} 0x656c60
				实例对象相等
				--- PASS: TestSingleton (0.00s)
				PASS
				ok      command-line-arguments  0.511s
	 ---------------------------------------------------------------

		1，测试单个文件，一定要带上被测试的原文件
    	go test -v  wechat_test.go  wechat.go

		2，测试单个方法
    	go test -v -test.run TestRefreshAccessToken
*/
