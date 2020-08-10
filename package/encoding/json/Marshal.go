/*
	func Marshal
	func Marshal(v interface{}) ([]byte, error)

	Marshal函数返回v的json编码。

	Marshal函数会递归的处理值。如果一个值实现了Marshaler接口切非nil指针，
	会调用其MarshalJSON方法来生成json编码。nil指针异常并不是严格必需的，但会模拟与UnmarshalJSON的行为类似的必需的异常。

	否则，Marshal函数使用下面的基于类型的默认编码格式：

	布尔类型编码为json布尔类型。

	浮点数、整数和Number类型的值编码为json数字类型。

	字符串编码为json字符串。角括号"<"和">"会转义为"\u003c"和"\u003e"以避免某些浏览器吧json输出错误理解为HTML。
	基于同样的原因，"&"转义为"\u0026"。

	数组和切片类型的值编码为json数组，但[]byte编码为base64编码字符串，nil切片编码为null。

	结构体的值编码为json对象。每一个导出字段变成该对象的一个成员，除非：

	- 字段的标签是"-"
	- 字段是空值，而其标签指定了omitempty选项
	空值是false、0、""、nil指针、nil接口、长度为0的数组、切片、映射。对象默认键字符串是结构体的字段名，
	但可以在结构体字段的标签里指定。结构体标签值里的"json"键为键名，后跟可选的逗号和选项，举例如下：

	// 字段被本包忽略
	Field int `json:"-"`
	// 字段在json里的键为"myName"
	Field int `json:"myName"`
	// 字段在json里的键为"myName"且如果字段为空值将在对象中省略掉
	Field int `json:"myName,omitempty"`
	// 字段在json里的键为"Field"（默认值），但如果字段为空值会跳过；注意前导的逗号
	Field int `json:",omitempty"`
	"string"选项标记一个字段在编码json时应编码为字符串。它只适用于字符串、浮点数、整数类型的字段。
	这个额外水平的编码选项有时候会用于和javascript程序交互：

	Int64String int64 `json:",string"`
	如果键名是只含有unicode字符、数字、美元符号、百分号、连字符、下划线和斜杠的非空字符串，将使用它代替字段名。

	匿名的结构体字段一般序列化为他们内部的导出字段就好像位于外层结构体中一样。如果一个匿名结构体字段的标签给其提供了键名，
	则会使用键名代替字段名，而不视为匿名。

	Go结构体字段的可视性规则用于供json决定那个字段应该序列化或反序列化时是经过修正了的。
	如果同一层次有多个（匿名）字段且该层次是最小嵌套的（嵌套层次则使用默认go规则），会应用如下额外规则：

	1）json标签为"-"的匿名字段强行忽略，不作考虑；

	2）json标签提供了键名的匿名字段，视为非匿名字段；

	3）其余字段中如果只有一个匿名字段，则使用该字段；

	4）其余字段中如果有多个匿名字段，但压平后不会出现冲突，所有匿名字段压平；

	5）其余字段中如果有多个匿名字段，但压平后出现冲突，全部忽略，不产生错误。

	对匿名结构体字段的管理是从go1.1开始的，在之前的版本，匿名字段会直接忽略掉。

	映射类型的值编码为json对象。映射的键必须是字符串，对象的键直接使用映射的键。

	指针类型的值编码为其指向的值（的json编码）。nil指针编码为null。

	接口类型的值编码为接口内保持的具体类型的值（的json编码）。nil接口编码为null。

	通道、复数、函数类型的值不能编码进json。尝试编码它们会导致Marshal函数返回UnsupportedTypeError。

	Json不能表示循环的数据结构，将一个循环的结构提供给Marshal函数会导致无休止的循环。
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
    ID     int
    Name   string
    Colors []string
}


func main(){
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b,err:=json.Marshal(group)
	if err !=nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

/*
PS D:\golang\github\golang_project\package\encoding\json> go run Marshal.go
{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
*/