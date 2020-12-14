package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    //创建一个新文件，写入内容 5 句 “http://c.biancheng.net/golang/”   
    filePath := "D:/goLang/github/code/golang.txt"
    file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("文件打开失败", err)
    }
    //及时关闭file句柄
    defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	/*
	func NewWriter 
	func NewWriter(w io.Writer) *Writer
	NewWriter创建一个具有默认大小缓冲、写入w的*Writer。
	*/
    write := bufio.NewWriter(file)
    for i := 0; i < 5; i++ {
        write.WriteString("http://c.biancheng.net/golang/ \n")
    }
    // Flush将缓存的文件真正写入到文件中
    write.Flush()
}