package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	// \d 是数字
	reQQEmail = `(\d+)@qq.com`
)

//爬邮箱  函数
func GetEmail() {
	// 1. 去网站拿数据
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	if err != nil {
		fmt.Println(err, "http.Get url")
	}
	defer resp.Body.Close()
	// 2. 读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err, "ioutil.ReadAll")
	}
	//fmt.Printf("%T",pageBytes) pageBytes  8位无符号整型 byte 字节型 表示单个字符， uint8 长度是一个字节
	// 将字节转为字符串
	pagestr := string(pageBytes) //
	//fmt.Println(pagestr)
	// 过滤数据
	re := regexp.MustCompile(reQQEmail) // 解析并返回一个正则表达式。如果成功返回，该Regexp就可用于匹配文本。
	// -1 代表取全部记录
	results := re.FindAllStringSubmatch(pagestr, -1)
	fmt.Println(results)

}

func main() {
	GetEmail()
}
