package main 

import (
	"fmt"
)

// 定义一个接口
type Person interface {
	Say()
}

// 定义一个类型
type Zt struct {
	name string 
}

// 实现接口的方法
func (zt *Zt)Say(){
	fmt.Println("我是 zt类型的方法，我是",zt.name)
}

func test(p Person){
	// 如果要使用它的实现实例zt的name属性，那么就需要类型断言
	name:=p.(*Zt).name
	fmt.Println(name)
}



func main (){
  //  实例化
  zt:=zt{
	  name:"taotao",
  }
  test(&zt)
}


/*
	1.类型断言就是将接口类型的值(x)，转换成类型(T)。格式为：x.(T);
	2.类型断言的必要条件就是x是接口类型，非接口类型的x不能做类型断言；
	3.T可以是非接口类型，如想断言合法，则T必须实现x的接口；
	4.T也可以是接口，则x的动态类型也应该实现接口T；
	5.类型断言如果非法，运行失败时会导致错误，为了避免这种错误，应该总是使用下面的方式来进行类型断言
	6.接口变量的类型也可以使用一种特殊形式的 swtich 来检测:
	
			func classifier(items ...interface{}) {
			for i, x := range items {
				switch x.(type) {
				case bool:
					fmt.Printf("Param #%d is a bool\n", i)
				case float64:
					fmt.Printf("Param #%d is a float64\n", i)
				case int, int64:
					fmt.Printf("Param #%d is a int\n", i)
				case nil:
					fmt.Printf("Param #%d is a nil\n", i)
				case string:
					fmt.Printf("Param #%d is a string\n", i)
				default:
					fmt.Printf("Param #%d is unknown\n", i)
				}
			}
		}
*/