package main

import (
	"log"
	"runtime"
	"time"
)

type Road int

func findRoad(r *Road) {
	log.Println("road:", *r)
}

func entry() {
	var rd Road = Road(999)
	r := &rd
	runtime.SetFinalizer(r, findRoad)
}

func main() {
	entry()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		runtime.GC()
	}
}

/**
PS D:\goLang\github\golang_project\Go语言结构体\Go语言垃圾回收和SetFinalizer> go run 1.go
2020/12/08 10:35:47 road: 999
*/
