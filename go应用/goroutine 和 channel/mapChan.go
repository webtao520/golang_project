// 创建一个mapChan，最多可以存放10个map[string]string的key-val，演示写入和读取
package main

import (
	"fmt"
)

func main (){
	var mapChan chan map[string]string
	mapChan =make(chan map[string]string,10)
	m1:=make(map[string]string,20)
	m1["city1"]="北京"
	m1["city2"]="上海"
	m2:=make(map[string]string,20)
	m2["hero1"]="松江"
	m2["hero2"]="武松"
	// ....
	mapChan <-m1
	mapChan<-m2

	
	fmt.Println(<-mapChan)
		
	fmt.Println(<-mapChan)
}


/*
map[city1:北京 city2:上海]
map[hero1:松江 hero2:武松]
*/