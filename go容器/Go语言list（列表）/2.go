package main

import (
	"container/list"
	"fmt"
)

func main() {
    l := list.New()

    // 尾部添加
    l.PushBack("canon")

    // 头部添加
    l.PushFront(67)

    // 尾部添加后保存元素句柄
    element := l.PushBack("fist")

    // 在fist之后添加high
    l.InsertAfter("high", element)

    // 在fist之前添加noon
    l.InsertBefore("noon", element)

    // 使用 移除尾部元素
	// l.Remove(element)
	

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

/**
PS D:\goLang\github\golang_project\go容器\Go语言list（列表）> go run 2.go
67
canon
noon
fist
high
*/