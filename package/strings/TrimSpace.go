/*
	func TrimSpace
	func TrimSpace(s string) string
	返回将s前后端所有空白（unicode.IsSpace指定）都去掉的字符串。
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n")) // a lone gopher
}
