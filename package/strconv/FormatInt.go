/*
	func FormatInt
	func FormatInt(i int64, base int) string
	返回i的base进制的字符串表示。base 必须在2到36之间，
	结果中会使用小写字母'a'到'z'表示大于10的数字。
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := strconv.FormatInt(50, 10)
	fmt.Println(str)
}
