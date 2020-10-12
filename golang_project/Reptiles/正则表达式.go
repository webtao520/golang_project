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
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}

func main() {

}
