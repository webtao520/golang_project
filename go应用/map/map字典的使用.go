/*
	map是一组键和值的组合。在map内是无序的。你可以随时加入或删除一个键及键对应的值。
	想查看所有键值，使用range遍历好了。
	想找到唯一的键值，那要使用键来得到值。
	map的创建必须使用make

	m := make(map [string]int)
	我们可以把这个map当作是一组人的名字和年龄，然后这样写入键值。


	m := make(map [string]int)
		m["Anna"] = 16
		m["Tom"] = 15
		m["Smith"] = 17
	直接打印map


	fmt. Println(m)
	运行结果


	map [Anna:16 Tom:15 Smith:17]
	如果需要挨个分析处理，那么需要遍历，并且判断后再加上处理语句。比如，我们给值为 17 的加一个特殊的显示标记“* —> ”


	//遍历，分别处理
		for k, v := range m {
			if v == 17 {fmt.Print("* --> ")}
			fmt.Println(k, v)
		}
	当然，你也可以直接对你感兴趣的map键调用。



	m["Anna"] = 17
	fmt.Println(m)
	这个时候的打印输出就会发现，Anna的值也是 17 了。
	如果，你只是想判断一下，map内是否已经有某个键了，可以这样写



	kv , ok := m["anna"]
		fmt.Println(kv, ok)
		kv1, ok1 := m["Anna"]
		fmt.Println(kv1, ok1)
	查看输出的结果。比较就会发现，当键不存在时，返回的值为零值，而同时我们得到了一个 false 的检测结果。我们可以使用这个方法来判断map中某个键是否存在。
	我们还需要知道，如何删除一个已经存在的元素。很简单的，只要直接 delete就行。像这样


	delete (m, "Anna")
	fmt.Println(m)
	观察打印结果，发现 Anna 已经不见了。

	完整代码 见下面

*/

/*
func main() {
	m := make(map[string]int)
	m["Anna"] = 16
	m["Tome"] = 15
	m["Smith"] = 17
	//fmt.Println(m) // map[Anna:16 Smith:17 Tome:15]

	//遍历，分别处理
	for k, v := range m {
		if v == 17 {
			fmt.Print("* ---> ")
			fmt.Println(k, v) // * ---> Smith 17
		}
	}
	m["Anna"] = 17
	fmt.Println(m) //map[Anna:17 Smith:17 Tome:15]

	kv, ok := m["anna"]
	fmt.Println(kv, ok) // 0 false
	kv1, ok1 := m["Anna"]
	fmt.Println(kv1, ok1) // 17 true

	//删除值
	delete(m, "Anna")
	fmt.Println(m) // map[Smith:17 Tome:15]

}
*/
/**
  ====================运行结果====================

PS D:\goLang\github\golang_project\go应用\map> go run .\map字典的使用.go
		* ---> Smith 17
		map[Anna:17 Smith:17 Tome:15]
		0 false
		17 true
		map[Smith:17 Tome:15]

*/

//  ==================  下面来展示一个复杂一点的例子。map元素值的类型是一个结构体。   ==================

package main

import "fmt"

type Vertex struct {
	Lat  float64
	Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Bei Jing"] = Vertex{39.26, 115.25}
	fmt.Println(m) // map[Bei Jing:{39.26 115.25} Bell Labs:{40.68433 -74.39967}]
	fmt.Println("--------I'm just a line.---------------")
	for lat, _ := range m {
		fmt.Println(lat, "：", m[lat].Lat, m[lat].Long)
		//Bell Labs ： 40.68433 -74.39967  Bei Jing ： 39.26 115.25
	}
}

/*
	  ====================运行结果====================

	PS D:\goLang\github\golang_project\go应用\map> go run .\map字典的使用.go
				map[Bei Jing:{39.26 115.25} Bell Labs:{40.68433 -74.39967}]
				--------I'm just a line.---------------
				Bell Labs ： 40.68433 -74.39967
				Bei Jing ： 39.26 115.25

*/
