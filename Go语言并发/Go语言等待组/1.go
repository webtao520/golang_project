package  main 


import (
	"fmt"
	"net/http"
	"sync"
)

func main(){
	//  声明一个等待组
	var wg  sync.WaitGroup

	// 准备一系列的网站地址
	var urls= []string{
		"http://www.github.com/",
        "https://www.qiniu.com/",
        "https://www.golangtc.com/",
	}

	// 遍历这些地址
	for _,url:=range urls {
		// 每一个任务开始时，将等待组增加 1 
		wg.Add(1)

		//  开启一个并发

		go func (url string){
			// 使用defer，表示函数完成时将等待组值减1
			defer wg.Done()  // wg,Add(-1)

			// 使用http 访问提供的地址
			_,err:=http.Get(url)

			// 访问完成后，打印地址和可能发生的错误
			fmt.Println(url,err)

			// 通过参数传递url地址
		}(url)
	}

	//  等待所有的任务完成
	wg.Wait()

	fmt.Println("over")
	
}

/**
PS D:\goLang\github\golang_project\Go语言并发\Go语言等待组> go run 1.go
https://www.qiniu.com/ <nil>
http://www.github.com/ <nil>
https://www.golangtc.com/ <nil>
over

代码说明如下：
第 12 行，声明一个等待组，对一组等待任务只需要一个等待组，而不需要每一个任务都使用一个等待组。
第 15 行，准备一系列可访问的网站地址的字符串切片。
第 22 行，遍历这些字符串切片。
第 25 行，将等待组的计数器加1，也就是每一个任务加 1。
第 28 行，将一个匿名函数开启并发。
第 31 行，在匿名函数结束时会执行这一句以表示任务完成。wg.Done() 方法等效于执行 wg.Add(-1)。
第 34 行，使用 http 包提供的 Get() 函数对 url 进行访问，Get() 函数会一直阻塞直到网站响应或者超时。
第 37 行，在网站响应和超时后，打印这个网站的地址和可能发生的错误。
第 40 行，这里将 url 通过 goroutine 的参数进行传递，是为了避免 url 变量通过闭包放入匿名函数后又被修改的问题。
第 44 行，等待所有的网站都响应或者超时后，任务完成，Wait 就会停止阻塞。
*/