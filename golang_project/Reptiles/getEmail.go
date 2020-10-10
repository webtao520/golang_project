package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//  获取qq邮箱
  var  (
	  // \d 是数字
	  reQQEmail = `(\d+)@qq.com`
  )

  // 爬邮箱
  func GetEmail(){
	  //  1. 去网站拿数据
	  resp,err:=http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	  HandleError(err, "http.Get url")
	  defer resp.Body.Close() // 程序在使用完回复后必须关闭回复的主体。
	  // 2. 读取页面内容
	  pageBytes, err := ioutil.ReadAll(resp.Body)
	  HandleError(err, "ioutil.ReadAll")
	  //fmt.Printf("%T",pageBytes) // []uint8  字节型(字符类型)：byte（uint8别名）
	  // 字节转字符串
	   pageStr:=string(pageBytes)
	   //fmt.Println(pageStr)
	   // 过滤数据 过滤qq邮箱
	   re:=regexp.MustCompile(reQQEmail)
		 // -1代表取全部
		 // Find返回一个保管正则表达式re在b中的所有不重叠的匹配结果
		 //及其对应的（可能有的）分组匹配的结果的[][]string切片。
		 //如果没有匹配到，会返回nil。
		 // func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
		 results := re.FindAllStringSubmatch(pageStr, -1)
		 //fmt.Println(results)
		     // 遍历结果
			 for _, result := range results {
				fmt.Println("email:", result[0])
				fmt.Println("qq:", result[1])
			}
	  
  }

  // 处理异常
  func HandleError (err error, why string){
        if err !=nil {
			fmt.Println(why,err)
		}
		return
  }

 func main (){
	GetEmail()
 }