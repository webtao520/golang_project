package models

// 实例化models
var Cache *localCache = NewCache()

type localCache struct {
	data map[string]interface{}
}

func NewCache() *localCache {
	c := &localCache{}
	c.data = make(map[string]interface{})
	return c
}

// 获取缓存
func (this *localCache) Get(key string) interface{} {
	// 类型断言
    if v,ok:=this.data[key];ok {
       return v
	}
	return ""
}

//  写入cache
func (this *localCache) Put(key string,val interface{}) error {
	this.data[key]=val 
	return nil
}

// 判断缓存是否存在
func (this *localCache) IsExist(key string) bool {
	if _,ok:=this.data[key];ok{
		return true
	}
	return false
} 

// 删除缓存
func (this *localCache) Delete(key string) error{
	delete(this.data,key)
	return nil
}

