package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	ret := Split("babcbef", "b")         // 程序输出的结果
	want := []string{"", "a", "c", "ef"} //  期望的结果
	if !reflect.DeepEqual(ret, want) {   // 比较两个值是否相等
		t.Errorf("want:%v but got:%v\n", want, ret)
	}
}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}

// 还可以在go test命令后添加-run参数，它对应一个正则表达式，只有函数名匹配上的测试函数才会被go test命令执行

/*
PS D:\goLang\github\golang_project\单元测试\split_string> go test -v
=== RUN   TestSplit
--- PASS: TestSplit (0.00s)
PASS
ok      _/D_/goLang/github/golang_project/单元测试/split_string 0.187s
*/
