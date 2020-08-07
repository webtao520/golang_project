/*
		func Replace 
		func Replace(s, old, new string, n int) string
		返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串
*/

package  main 

import (
	"fmt"
	"strings"
)

func  main (){
	fmt.Println(strings.Replace("oink oink oink", "k", "zt", 3)) // oinzt oinzt oinzt
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))	//moo moo moo
}


/*
		##################### 运行结果 ##########################

		PS D:\goLang\github\golang_project\package\strings> go run .\Replace.go
		oinzt oinzt oinzt
		moo moo moo

*/