/*
	func QueryEscape
		func QueryEscape(s string) string
		QueryEscape函数对s进行转码使之可以安全的用在URL查询里。

	url 编码
	   QueryEscape 将字符地址转为安全地址

	QueryUnescape 还原QueryEscape 转码
	  func QueryUnescape(s string) (string, error)
*/

package main

import (
	"fmt"
	"net/url"
)

func main() {
	p := url.QueryEscape("/user?id=12")
	fmt.Println(p) // %2Fuser%3Fid%3D12
	str, err := url.QueryUnescape(p)
	if err == nil {
		fmt.Println(str)
	}
}

/*
=================== 运行结果 =======================
	PS D:\goLang\github\golang_project\package\url> go run .\QueryEscape.go
		%2Fuser%3Fid%3D12
		/user?id=12
*/
