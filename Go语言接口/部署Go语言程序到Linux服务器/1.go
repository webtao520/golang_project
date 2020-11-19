package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "C语言中文网")
}
func main() {
	http.HandleFunc("/", sayHello)
	log.Println("启动成功，可以通过 localhost:9000 访问")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("List 9000")
	}
}
