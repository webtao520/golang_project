/*
			package md5
			import "crypto/md5"
			md5包实现了MD5哈希算法

func Sum
		func Sum(data []byte) [Size]byte
		返回数据data的MD5校验和。


*/

package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("These pretzels are making me thirsty.")
	fmt.Printf("%x", md5.Sum(data)) // %x	表示为十六进制，使用a-f
	fmt.Println()
	fmt.Println(data)
}

/*
PS D:\golang\github\golang_project\package\md5> go run .\Sum.go

		b0804ec967f48520697662a204f5fe72
		[84 104 101 115 101 32 112 114 101 116 122 101 108 115 32 97 114 101 32 109 97 107 105 110 103 32 109 101 32 116 104 105 114 115 116 121 46]

*/
