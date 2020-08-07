/*
	func Trim
	func Trim(s string, cutset string) string
	返回将s前后端所有cutset包含的utf-8码值都去掉的字符串。

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))
}

/*
	################# 运行结果 ############################
		PS D:\goLang\github\golang_project\package\strings> go run .\Trim.go
		["Achtung! Achtung"]
*/
