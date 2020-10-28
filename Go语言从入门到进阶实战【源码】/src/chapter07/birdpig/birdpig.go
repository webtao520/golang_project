package main

import "fmt"

// 飞行物接口
type Flyer interface {
	Fly()
}

// 行走物接口
type Walker interface {
	Walk()
}

// 鸟类
type bird struct {
}

// 实现飞行接口
func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

// 实现行走接口
func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

// 猪
type pig struct {
}

// 实现行走接口
func (p *pig) Walk() {
	fmt.Println("pig: walk")
}

func main() {

	// 动物的名字到实例的映射
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}

	// 遍历映射
	for name, obj := range animals {

		// 断言对象是否为飞行物
		f, isFlyer := obj.(Flyer)
		// 断言对象是否为行走物
		w, isWalker := obj.(Walker)

		fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)

		// 飞行物调用飞行接口
		if isFlyer {
			f.Fly()
		}

		// 行走物调用行走接口
		if isWalker {
			w.Walk()
		}
	}

}
