// 字典结构
type  Dictionary  struct {
	data map[interface{}]interface{} // 键值都为interface{}类型
}

// 根据健获取值
func (d *Dictionary) Get(key  interface {}) interface{}{
   return d.data[key]
}

// 设置键值
func (d *Dictionary)Set(key interface{},value interface{}){
	 d.data[key]=value
}

/**
代码说明如下：
第 3 行，Dictionary 的内部实现是一个键值均为 interface{} 类型的 map，map 也具备与 Dictionary 一致的功能。
第 8 行，通过 map 直接获取值，如果键不存在，将返回 nil。
第 13 行，通过 map 设置键值。
*/