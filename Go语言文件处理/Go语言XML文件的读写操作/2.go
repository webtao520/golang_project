package main

import (
	"fmt"
	"os"
	"encoding/xml"
)

type Website  struct {
	Name string `xml:"name,attr"`
	Url string
	Course []string
}

func main(){
	// 打开xml 文件
	file, err := os.Open("./info.xml")
	if err !=nil{
		fmt.Printf("文件打开失败：%v", err)
        return
	}
	defer file.Close()

	info:=Website{}
	    //创建 xml 解码器
		decoder := xml.NewDecoder(file)
		err = decoder.Decode(&info)
		if err != nil {
			fmt.Printf("解码失败：%v", err)
			return
		} else {
			fmt.Println("解码成功")
			fmt.Println(info)
		}
}

/**
PS D:\goLang\github\golang_project\Go语言文件处理\Go语言XML文件的读写操作> go run 2.go
解码成功
{C语言中文网 http://c.biancheng.net/golang/ [Go语言入门教程 Golang入门教程]}
*/