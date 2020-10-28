package cls1

import (
	"chapter08/clsfactory/base"
	"fmt"
)

// 类1
type Class1 struct {
}

// 实现Class接口
func (c *Class1) Do() {
	fmt.Println("Class1")
}

func init() {

	// 在启动时注册类1工厂
	base.Register("Class1", func() base.Class {
		return new(Class1)
	})
}
