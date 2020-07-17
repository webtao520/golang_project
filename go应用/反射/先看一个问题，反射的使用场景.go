package main 

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json:"monsterName"`
	Age int `json:"monsterAge"`
	Sal float64 `json:"monsterSal"`
	Sex  string `json:"monsterSex"` 
}

func main (){
	 m:=Monster {
		 Name:"age",
		 Age:30,
		 Sal:34.34,
		 Sex:"female",
	 }
	 data,_:=json.Marshal(m)
	 fmt.Println("json result:",string(data))
}

/*
PS D:\goLang\github\golang_project\go应用\反射> go run .\先看一个问题，反射的使用场景.go
json result: {"monsterName":"age","monsterAge":30,"monsterSal":34.34,"monsterSex":"female"}
*/