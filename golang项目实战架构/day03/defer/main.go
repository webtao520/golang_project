package main 

import (
	"fmt"
)

//defer
func calc(index string,a,b int) int{
   ret :=a+b 
   fmt.Println(index,a,b,ret)
   return ret
}

func main(){
	a:=1
	b:=2
	defer calc("1", a, calc("10", a, b))
	a= 0 
	defer calc("2", a, calc("20", a, b))
	b = 1
}

/**
PS D:\goLang\github\golang_project\golang项目实战架构\day03\defer> go run .\main.go
// 1. a:=1
// 2. b:=2
// 3. defer calc("1", 1, calc("10", 1, 2))
// 4. calc("10", 1, 2) // "10" 1 2 3
// 5. defer calc("1", 1, 3)
// 6. a = 0
// 7. defer calc("2", 0, calc("20", 0, 2))
// 8. calc("20", 0, 2) // "20" 0 2 2
// 9. defer calc("2", 0, 2)
// 10. b = 1
// calc("2", 0, 2) // "2" 0 2 2
// calc("1", 1, 3) // "1" 1 3 4

// 最终的答案：
// "10" 1 2 3
// "20" 0 2 2
//  "2" 0 2 2
// "1" 0 3 3
*/