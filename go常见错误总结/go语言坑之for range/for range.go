/**
go只提供了一种循环方式，即for循环，在使用时可以像c那样使用，也可以通过for range方式遍历容器类型如数组、切片和映射。但是在使用for
range时，如果使用不当，就会出现一些问题，导致程序运行行为不如预期。比如，下面的示例程序将遍历一个切片，并将切片的值当成映射的键和值存入，
切片类型是一个int型，映射的类型是键为int型，值为*int，即值是一个地址。
*/

/*
package main 

import (
	"fmt"
)

func main () {
slice :=[]int{0,1,2,3}
myMap:=make(map[int]*int)
for index,value:=range slice {
	myMap[index]=&value
}
fmt.Println(myMap)
fmt.Println("========== new map ===========")
pritMap(myMap)
}

func  pritMap(myMap map[int]*int)  {
	 for key,value:=range myMap {
		 fmt.Printf("map[%v]=%v\n",key,*value)
	 }
}
*/



/**
PS D:\goLang\github\golang_project\go常见错误总结\go语言坑之for range> go run '.\for range.go'
					==========new map===========
					map[0]=3
					map[1]=3
					map[2]=3
					map[3]=3

但是由输出可以知道，映射的值都相同且都是3。其实可以猜测映射的值都是同一个地址，遍历到切片的最后一个元素3时，
将3写入了该地址，所以导致映射所有值都相同。其实真实原因也是如此，因为for
range创建了每个元素的副本，而不是直接返回每个元素的引用，如果使用该值变量的地址作为指向每个元素的指针，
就会导致错误，在迭代时，返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以值的地址总是相同的，导致结果不如预期。

修正后程序如下：

*/

package  main 

import (
	"fmt"
)

func main (){
	slice:=[]int{0,1,2,3}
	myMap:=make(map[int]*int)
	
	for index,value:=range slice {
		num:=value // 开辟新的内存空间 
		myMap[index]=&num
	}
	fmt.Println("=====new map=====")
    prtMap(myMap)
}


func prtMap(myMap map[int]*int) {
    for key, value := range myMap {
        fmt.Printf("map[%v]=%v\n", key, *value)
    }
}

/*
PS D:\goLang\github\golang_project\go常见错误总结\go语言坑之for range> go run '.\for range.go'
				=====new map=====
				map[0]=0
				map[1]=1
				map[2]=2
				map[3]=3
*/

