package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币

写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数



//  uint8 byte 和  int32 rune 区别
// https://www.cnblogs.com/wjaaron/p/11485685.html



PS D:\goLang\github\golang_project\golang项目实战架构\day03\homework> go run main.go
剩下： 10
Matthew:1
Augustus:12
Peter:2
Adriano:5
Aaron:3
Elizabeth:4
Heidi:5
Emilie:6
Giana:2

*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution=make(map[string]int,len(users))

)

func dispatchCoin() (left int) {
	// 1. 依次拿到每个人的名字
	for _,name:=range  users {
	// 2. 拿到一个人名根据分金币的规则去分金币,  字符串底层是 字节数组
	for _,c:=range name { // int32 rune  
		switch c {
		case 'e','E':
		  // 满足这个条件分1枚金币   (使用range，其实是使用rune类型来编码的，rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。)
		  distribution[name]++
		  coins--
		case 'i', 'I':
			// 满足这个条件分2枚金币
			distribution[name] += 2
			coins -= 2
		case 'o', 'O':
			// 满足这个条件分3枚金币
			distribution[name] += 3
			coins -= 3
		case 'u', 'U':
			// 满足这个条件分4枚金币
			distribution[name] += 4
			coins -= 4	
		}	
		left = coins
	}
	// fmt.Println(name)
	}
	return
}

func main(){
	left:=dispatchCoin()
	fmt.Println("剩下：", left)
	for k,v:=range distribution{
		fmt.Printf("%s:%d\n",k,v)
	}
}

