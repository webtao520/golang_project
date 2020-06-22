package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// 正则表达式
var (
	// 匹配电子邮箱  \w 匹配数字，字母，下划线任意字符，	匹配除换行符 \n 之外的任何单字符。要匹配 . ，请使用 \. 。
	//+ 匹配前面的子表达式一次或多次。要匹配 + 字符，请使用 \+。
	reEmail = `\w+@\w+\.\w+`
	// s?  匹配前面的子表达式零次或一次，或指明一个非贪婪限定符
	reLinke  = `href="(https?://[\s\S]+?)"`
	rePhone  = `1[3456789]\d\s\d{4}\s?\d{4}` // \d 匹配一个数字字符  等价于[0-9]
	reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
	// 指明两项之间的一个选择，要匹配|  \|
	reImg = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// email
func GetEmail2(url string) {
	pageStr := GetPageStr(url)
	//fmt.Println(pageStr)
	re := regexp.MustCompile(reEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	// fmt.Printf("%T",results)
	for _, result := range results {
		fmt.Println(result)
	}
}

// 根据url 获取网页内容
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}

// 连接
func GetLink(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reLinke)
	results := re.FindAllStringSubmatch(pageStr, -1)
	//    for i:=0;i<len(results);i++ {
	// 	   fmt.Println(results[i])
	//    }
	for k, result := range results {
		if k == 0 {
			fmt.Println(result[k])
		}
		break
	}

}

// 获取身份证号
func GetIdCard(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reIdcard)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Println(results)
	for _, result := range results {
		fmt.Println(result)
	}
}

// 获取手机号
func GetPhone(url string) {
	pageStr := GetPageStr(url)
	//fmt.Println(pageStr)
	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

func GetImg(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[0])
	}
}

func main() {
	// 1.抽取的爬邮箱
	// GetEmail2("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	// 2.爬链接
	//GetLink("http://www.baidu.com/s?wd=%E8%B4%B4%E5%90%A7%20%E7%95%99%E4%B8%8B%E9%82%AE%E7%AE%B1&rsv_spt=1&rsv_iqid=0x98ace53400003985&issp=1&f=8&rsv_bp=1&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_dl=ib&rsv_sug2=0&inputT=5197&rsv_sug4=6345")
	// 3.爬手机号
	//GetPhone("https://www.zhaohaowang.com/")
	// 4.爬身份证号
	// GetIdCard("https://henan.qq.com/a/20171107/069413.htm")
	// 5.爬图片
	GetImg("http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3")
}
