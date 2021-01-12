package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")

	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(jpgFunc("呵呵.jpg"))
	fmt.Println(txtFunc("test")) //test.txt
	fmt.Println(txtFunc("呵呵.txt"))
}
