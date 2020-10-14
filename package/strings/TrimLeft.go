/*
func TrimLeft
func TrimLeft(s string, cutset string) string
返回将s前端所有cutset包含的utf-8码值都去掉的字符串。
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "golang php"
	fmt.Printf("%v", strings.TrimLeft(str, "go"))
}

/**
PS D:\golang\github\golang_project\package\strings> go run TrimLeft.go
lang php
*/
