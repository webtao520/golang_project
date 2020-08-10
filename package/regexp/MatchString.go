/*
	func MatchString
	func MatchString(pattern string, s string) (matched bool, err error)

	MatchString类似Match，但匹配对象是字符串。
*/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.MatchString("foo.*", "seafood")
	fmt.Println(matched, err)
}

/*
	PS D:\golang\github\golang_project\package\regexp> go run .\MatchString.go
	true <nil>
*/
