package main

import "fmt"

// 定义飞行动物接口
type Flyer interface {
    Fly()
}

// 定义行走动物接口
type Walker interface {
    Walk()
}

// 定义鸟类
type bird struct {
}

// 实现飞行动物接口
func (b *bird) Fly() {
    fmt.Println("bird: fly")
}

// 为鸟添加Walk()方法, 实现行走动物接口
func (b *bird) Walk() {
    fmt.Println("bird: walk")
}

// 定义猪
type pig struct {
}

// 为猪添加Walk()方法, 实现行走动物接口
func (p *pig) Walk() {
    fmt.Println("pig: walk")
}

func main() {

// 创建动物的名字到实例的映射
    animals := map[string]interface{}{
        "bird": new(bird),
        "pig":  new(pig),
    }

    // 遍历映射
    for name, obj := range animals {

        // 判断对象是否为飞行动物
        f, isFlyer := obj.(Flyer)
        // 判断对象是否为行走动物
        w, isWalker := obj.(Walker)

        fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)

        // 如果是飞行动物则调用飞行动物接口
        if isFlyer {
            f.Fly()
        }

        // 如果是行走动物则调用行走动物接口
        if isWalker {
            w.Walk()
        }
    }
}