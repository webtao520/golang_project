// 清空所有的数据
func (d *Dictionary) Clear (){
	d.data=make(map[interface{}] interface{})
}

//创建一个字典
fun NewDictionary () *Dictionary{
	d:=&Dictionary{}

	// 初始化map
	d.Clear()
	return d
}