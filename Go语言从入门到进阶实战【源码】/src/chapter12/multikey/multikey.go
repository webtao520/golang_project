package main

import (
	"fmt"
)

type Profile struct {
	Name    string
	Age     int
	Married bool
}

// 查询键
type queryKey struct {
	Name string
	Age  int
}

// 创建查询键到数据的映射
var mapper = make(map[interface{}]*Profile)

// 构建查询索引
func buildIndex(list []*Profile) {

	// 遍历所有数据
	for _, profile := range list {

		// 构建查询键
		key := queryKey{
			Name: profile.Name,
			Age:  profile.Age,
		}

		// 保存查询键
		mapper[key] = profile
	}
}

// 根据条件查询数据
func queryData(name string, age int) {

	// 根据查询条件构建查询键
	key := queryKey{name, age}

	// 根据键值查询数据
	result, ok := mapper[key]

	// 找到数据打印出来
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("no found")
	}
}

func main() {

	list := []*Profile{
		{Name: "张三", Age: 30, Married: true},
		{Name: "李四", Age: 21},
		{Name: "王麻子", Age: 21},
	}

	buildIndex(list)

	queryData("张三", 30)
}
