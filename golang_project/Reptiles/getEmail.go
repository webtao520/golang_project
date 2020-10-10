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
	  // re := regexp.MustCompile(reStr)，传入正则表达式，得到正则表达式对象
	 //	ret := re.FindAllStringSubmatch(srcStr,-1)：用正则对象，获取页面页面，srcStr是页面内容，-1代表取全部
	   re:=regexp.MustCompile(reQQEmail)
		 // -1代表取全部
		 results := re.FindAllStringSubmatch(pageStr, -1)
		 //fmt.Println(results)
		 mapEmali:=[]string{}
		 mapQQ:=[]string{}
		     // 遍历结果
			 for _, result := range results {
				//fmt.Println("email:", result[0])
				mapEmali = append(mapEmali, result[0])  // 向切片中添加元素
				mapQQ = append(mapQQ, result[1]) 
				//fmt.Println("qq:", result[1])
			}
			resEmali:=removeRepByMap(mapEmali)
			resQQ:=removeRepByMap(mapQQ)
			fmt.Println(resEmali,resQQ)
  }

  func removeRepByMap(slc []string) []string {
    result := []string{}        
    tempMap := map[string]byte{}
    for _, e := range slc {
		l := len(tempMap)
        tempMap[e] = 0 
		
        if len(tempMap) != l { 
            result = append(result, e) 
        }
    }
    return result
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