package main 


import (
	"fmt"
	"io/ioutil"
)

func readFromFileByIoutil() {
	content,err:=ioutil.ReadFile("./main.go")
	if err !=nil{
	  fmt.Println("read file failed, err:", err)
	  return
	}
	fmt.Println(string(content))
}

// io/ioutil包的ReadFile方法能够读取完整的文件，只需要将文件名作为参数传入
func  main(){
	readFromFileByIoutil()
}


