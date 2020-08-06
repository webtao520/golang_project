package main 

import (
	"fmt"
)

func main () {
	str:=fmt.Sprintf("%[2]d%[1]d\n", 11, 22)
	fmt.Println(str)
}

/*
  ############################ 运行结果 #############################

PS D:\goLang\github\golang_project\package\fmt> go run .\Sprintf.go
2211
*/