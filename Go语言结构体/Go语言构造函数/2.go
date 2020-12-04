package main

import (
	"fmt"
)

type Cat struct {
	Color string
	Name  string
}

type BlackCat struct {
	Cat // 嵌入Cat ,类似派生
}

// “构造积累”
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// 构造子类
func NewBlackCat (color string) *BlackCat{
	cat:=&BlackCat{}
	cat.Color =color
	return cat

}

