/**
文档：https://studygolang.com/pkgdoc

API
re := regexp.MustCompile(reStr)，传入正则表达式，得到正则表达式对象
ret := re.FindAllStringSubmatch(srcStr,-1)：用正则对象，获取页面页面，srcStr是页面内容，-1代表取全部

爬邮箱
方法抽取
爬超链接
爬手机号
http://www.zhaohaowang.com/ 如果连接失效了自己找一个有手机号的就好了

爬身份证号
http://henan.qq.com/a/20171107/069413.htm 如果连接失效了自己找一个就好了
爬图片链接
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	// w代表大小写字母+数字+下划线
	reEmail = `\w+@\w+\.\w+`
	// s?有或者没有s
	// +代表出1次或多次
	//\s\S各种字符
	// +?代表贪婪模式
	reLinke  = `href="(https?://[\s\S]+?)"`
	rePhone  = `1[3456789]\d\s?\d{4}\s?\d{4}`
	reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
	reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// 根据 url 获取内容
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2. 读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	//  字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}

func main() {
	// 2.抽取的爬邮箱
	// GetEmail2("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	// 3.爬链接
	//GetLink("http://www.baidu.com/s?wd=%E8%B4%B4%E5%90%A7%20%E7%95%99%E4%B8%8B%E9%82%AE%E7%AE%B1&rsv_spt=1&rsv_iqid=0x98ace53400003985&issp=1&f=8&rsv_bp=1&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_dl=ib&rsv_sug2=0&inputT=5197&rsv_sug4=6345")
	// 4.爬手机号
	//GetPhone("https://www.zhaohaowang.com/")
	// 5.爬身份证号
	//GetIdCard("https://henan.qq.com/a/20171107/069413.htm")
	// 6.爬图片
	// GetImg("http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3")
}

/**
re := regexp.MustCompile(reStr)，传入正则表达式，得到正则表达式对象
ret := re.FindAllStringSubmatch(srcStr,-1)：用正则对象，获取页面页面，srcStr是页面内容，-1代表取全部
*/
// 获取身份证号
func GetIdCard(url string) {
	pageStr := GetPageStr(url) //  获取页面内容
	// 正则匹配
	re := regexp.MustCompile(reIdcard)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

// 爬链接
func GetLink(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reLinke)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[1])
	}
}

//爬手机号
func GetPhone(url string) {
	pageStr := GetPageStr(url)
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
