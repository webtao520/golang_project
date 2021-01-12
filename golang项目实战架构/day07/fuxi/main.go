package main 

import (
	"fmt"
	"encoding/json"
)

// 反射
type  student struct {
	 Name string `json:name`
	 Age int `json:age`
}



func main (){
   str:=  `{"name":"张三","age":9000}`
   var stu1 student
   json.Unmarshal([]byte(str),&stu1)
   fmt.Printf("%#v\n",stu1) // main.student{Name:"张三", Age:9000}
}
