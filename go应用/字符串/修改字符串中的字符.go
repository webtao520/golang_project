/**
	修改字符串中的字符(用 []rune)
*/

package main 

import (
	"fmt"
)

func main(){
	fmt.Println("=======byte======= uint8 =================")
	s:="Hello 世界"
	b:=[]byte(s)
	b[0]='E'
	fmt.Println(string(b)) // Eello 世界

	fmt.Println("=======rune======== int32 ================")
	b:=[]rune(s) 
	b[6] = '中'
	b[7] = '国'
	fmt.Println(s)
	fmt.Printf("%T\n",string(b))
}