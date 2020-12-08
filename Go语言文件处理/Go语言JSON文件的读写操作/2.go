package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func main() {
	filePtr, err := os.Open("./info.json")
	if err != nil {
		fmt.Println("文件打开失败 [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()
	var info []Website
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}

/**
解码成功
[{Golang http://c.biancheng.net/golang/ [http://c.biancheng.net/cplus/ http://c.biancheng.net/linux_tutorial/]}
{Java http://c.biancheng.net/java/ [http://c.biancheng.net/socket/ http://c.biancheng.net/python/]}]
*/
