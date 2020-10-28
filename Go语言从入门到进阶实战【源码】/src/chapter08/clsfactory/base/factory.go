package base

// 类接口
type Class interface {
	Do()
}

var (
	// 保存注册好的工厂
	factoryByName = make(map[string]func() Class)
)

// 注册一个类生成工厂
func Register(name string, factory func() Class) {
	factoryByName[name] = factory
}

// 根据名称创建对应的类
func Create(name string) Class {
	if f, ok := factoryByName[name]; ok {
		return f()
	} else {
		panic("name not found")
	}
}
