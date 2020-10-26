package main 

import (
	"fmt"
)

func main(){
	noteFre:=map[string]float32{
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
		"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440,
	}
	fmt.Println(noteFre)
}

/*
	PS D:\goLang\github\golang_project\go容器\Go语言map（映射）> go run 2.go
	map[A0:27.5 A4:440 B0:30.87 C0:16.35 D0:18.35 E0:20.6 F0:21.83 G0:24.5]
*/