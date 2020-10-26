package main 

import (
	"fmt"
)

func main(){
	scene := make(map[string]int)
	// 准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	delete(scene,"brazil")

	for k,v:=range scene{
		fmt.Println(k,v)
	}
}

/**
PS D:\goLang\github\golang_project\go容器\map元素的删除和清空> go run 1.go
route 66
china 960
*/