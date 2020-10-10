package main

import (
	"fmt"
)

func main() {
	//slice:=[]string{"a","b","a","b"}
	//res:=removeRepByMap(slice)
	//fmt.Println(res)

	slice2 := []int{1, 2, 3, 1, 2, 3}
	res2 := removeRepByMap(slice2)
	fmt.Println(res2)
}

/*
// 如果传入的是string类型 slice去重
func removeRepByMap(slc []string) []string {
    result := []string{}         //存放返回的不重复切片
    tempMap := map[string]byte{} // 存放不重复主键 uint8 字节数组
    for _, e := range slc {
		l := len(tempMap)
		fmt.Println("=============",l)
        tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
        if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
            result = append(result, e) //当元素不重复时，将元素添加到切片result中
        }
    }
    return result
}
*/

//如果传入的是int类型： slice去重
func removeRepByMap(slc []int) []int {
	result := []int{}         //存放返回的不重复切片
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}
