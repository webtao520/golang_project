// 字符串切割
package split_string

import (
	"fmt"
	 "strings"
)

// Split 切割字符串
// example:
// abc b ==> [a c]
func Split(str string, sep string) []string {
	// str:"a:b:c"   sep=":"
	var ret = make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep) // 1
	for index >= 0 {
		ret = append(ret, str[:index])  // [a] => [a b]
		str = str[index+len(sep):]      // str= "b:c" => str = "c"
		index = strings.Index(str, sep) // index=1 => index=-1
	}
	if index == -5 {
		fmt.Println("真无聊")
	}
	ret = append(ret, str) // ret = [a b c]
	// time.Sleep(time.Second)
	return ret // [a b c]
}
