/*
		func Unmarshal
		func Unmarshal(data []byte, v interface{}) error

		Unmarshal函数解析json编码的数据并将结果存入v指向的值。

		Unmarshal和Marshal做相反的操作，必要时申请映射、切片或指针，有如下的附加规则：

		要将json数据解码写入一个指针，Unmarshal函数首先处理json数据是json字面值null的情况。
		此时，函数将指针设为nil；否则，函数将json数据解码写入指针指向的值；如果指针本身是nil，
		函数会先申请一个值并使指针指向它。

		要将json数据解码写入一个结构体，函数会匹配输入对象的键和Marshal使用的键（结构体字段名或者它的标签指定的键名），
		优先选择精确的匹配，但也接受大小写不敏感的匹配。

		要将json数据解码写入一个接口类型值，函数会将数据解码为如下类型写入接口：

		Bool                   对应JSON布尔类型
		float64                对应JSON数字类型
		string                 对应JSON字符串类型
		[]interface{}          对应JSON数组
		map[string]interface{} 对应JSON对象
		nil                    对应JSON的null

		如果一个JSON值不匹配给出的目标类型，或者如果一个json数字写入目标类型时溢出，
		Unmarshal函数会跳过该字段并尽量完成其余的解码操作。如果没有出现更加严重的错误，
		本函数会返回一个描述第一个此类错误的详细信息的UnmarshalTypeError。

		JSON的null值解码为go的接口、指针、切片时会将它们设为nil，因为null在json里一般表示“不存在”。 
		解码json的null值到其他go类型时，不会造成任何改变，也不会产生错误。

		当解码字符串时，不合法的utf-8或utf-16代理（字符）对不视为错误，而是将非法字符替换为unicode字符

*/

package main 

import (
	"fmt"
	"encoding/json"
)

func main (){
     var jsonBlob= []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	 ]`)

	 type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}

/*
  ################### 运行结果 ###########################
	PS D:\golang\github\golang_project\package\encoding\json> go run .\Unmarshal.go
	[{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
*/