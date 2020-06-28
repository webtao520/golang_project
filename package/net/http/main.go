package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/*
   实现一个简单的http服务器作为web服务器
   ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，
   这表示采用包变量DefaultServeMux作为处理器。
   Handle和HandleFunc函数可以向DefaultServeMux添加处理器
*/

func sayHelloWord(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数
	fmt.Println(r.Form) // 在服务端打印请求参数
	fmt.Println("URL:", r.URL.Path)
	fmt.Println(r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println(k, ":", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "你好，golang") // 发送响应到客户端
}

/*
    要管理服务端的行为，可以一个自定义的server


 s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())
*/

func main() {
	http.HandleFunc("/", sayHelloWord) // HandleFunc函数可以向DefaultServeMux添加处理器
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
