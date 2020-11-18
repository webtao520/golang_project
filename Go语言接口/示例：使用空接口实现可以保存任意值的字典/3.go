// 遍历所有的键值, 如果回调返回值为false, 停止遍历
func (d *Dictionary) Visit(callback func(k,v interface{}) bool) {
	   if callback == nil {
		   return
	   }
	
	for k,v :=range d.data{
		if !callback(k,v) {
			return
		}
	}
}


/**
代码说明如下：
第 2 行，定义回调，类型为 func(k,v interface{})bool，意思是返回键值数据（k、v）。bool 表示遍历流程控制，返回 true 时继续遍历，返回 false 时终止遍历。
第 4 行，当 callback 为空时，退出遍历，避免后续代码访问空的 callback 而导致的崩溃。
第 8 行，遍历字典结构的 data 成员，也就是遍历 map 的所有元素。
第 9 行，根据 callback 的返回值，决定是否继续遍历。
*/