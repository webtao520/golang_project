/*
	介绍了管道类型的基本语法，通常，管道都是支持双向操作的：既可以往管道发送数据，也可以从管道接收数据。但在某些场景下，
	可能我们需要限制只能往管道发送数据，或者只能从管道接收数据，这个时候，就需要用到单向通道。

	不过，这里我们需要澄清一下，通道本身还是要支持读写的，如果某个通道只支持写入操作，那么即便数据写进去了，不能被读取也毫无意义，同理，
	如果某个通道只支持读取操作，不能写入数据，那么通道永远是空的，从一个空的通道读取数据会导致协程的阻塞，无法执行后续代码。
	
	因此，Go 语言支持的单向管道，实际上是在使用层面对通道进行限制，而不是语法层面：即我们在某个协程中只能对通道进行写入操作，
	而在另一个协程中只能对该通道进行读取操作。从这个层面来说，单向通道的作用是约束在生产协程中只能发送数据到通道，

	而在消费协程中只能从通道接收数据，从而让代码遵循「最小权限原则」，避免误操作和通道使用的混乱，让代码更加稳健。
	下面我们就来看看如何在 Go 协程之间实现单向通道的约束。
	当我们将一个通道类型变量传递到一个函数时（通常是在另外一个协程中执行），如果这个函数只能发送数据到通道，
	可以通过如下将其指定为单向只写通道（发送通道）：
	
	 func test(ch chan<- int)

	 上述代码限定在 test 函数中只能写入 int 类型数据到通道 ch。
	反过来，如果我们将一个通道类型变量传递到一个只允许从该通道读取数据的函数，可以通过如下方式将通道指定为单向只读通道（接收通道）

	func test(ch <-chan int)

	上述代码限定在 test 函数中只能从 ch 通道读取 int 类型数据。

	虽然我们也可以像声明正常通道类型那样声明单向通道，但我们一般不这么做，因为这样一来，就是从语法上限定通道的操作类型了，
	对于只读通道只能接收数据，对于只写通道只能发送数据：

	var ch1 chan int // 声明通道，默认是双向
	var ch2 chan<-int  // 单向 写入
	var ch3 <-chan int // 单向 // 读取

	单向通道的初始化和双向通道一样：

	ch1:=make(chan int)
	ch2:=make(chan<-int)
	ch3:=make(<-chan int)

	此外，我们还可以通过如下方式实现双向通道和单向通道的转化

	ch1:=make(chan int)
	ch2:=<-chan int(ch1)
	ch3:=chan<- int(ch1)

	基于双向通道 ch1，我们通过类型转化初始化了两个单向通道：单向只读的 ch2 和单向只写的 ch3。注意这个转化是不可逆的，双向通道可以转化为任意类型的单向通道，
	但单向通道不能转化为双向通道，读写通道之间也不能相互转化。
	实际上，我们在将双向通道传递到限定通道参数操作类型的函数时，就应用到了类型转化。
	我们可以通过单向通道来约束上篇教程的示例代码中子协程对通道的单向写入操作：

*/

package main 

import (
	"fmt"
	"time"
)

// 往通道中写入数据，发送数据 (单向通道)
func test(ch  chan<- int){
	for i:=0;i<100;i++{
		ch<- i
	}
	close(ch)
}


func  main()  {
   start :=time.Now()
   ch:=make(chan int,10)
   //ch:=make(chan<- int)  
   /*
   如果我们将 test 函数中的通道参数类型约束调整为 test(ch <-chan int)，编译代码就会报错：
	
   # command-line-arguments
	.\chan.go:74:15: syntax error: unexpected <-, expecting comma or )

	提示传入的通道是只读通道（receive-only channel），不能进行写入操作，此外，关闭通道函数 close 也不能作用到只读通道。

	如果将 main 函数中的通道初始化语句修改为 ch := make(chan<- int)，编译时也会报错：

	# command-line-arguments
	.\chan.go:84:11: invalid operation: range ch (receive from send-only type chan<- int)

	提示不能通过 range 语句从只写通道（send-only）中接收数据。

	我们也可以定义一个返回值类型为单向只读通道的函数，以便得到该返回值的代码只能从通道中接收数据

	我们也可以定义一个返回值类型为单向只读通道的函数，以便得到该返回值的代码只能从通道中接收数据：

    test2 函数
   */
   go test(ch)
   // go test2()
   for i:=range ch {
   fmt.Println("<-chan int :", i)	   
 }
  end:=time.Now()
  consume:=end.Sub(start).Seconds()
  fmt.Println("程序执行耗时(s)：", consume)
}

/*
	显然，合理使用单向通道，可以有效约束不同业务对通道的操作，避免越权使用和滥用，此外，也提高了代码的可读性，
	一看函数参数就可以判断出业务对通道的操作类型。
*/
func test2() <-chan int {
    ch := make(chan int, 20)
    for i := 0; i < 100; i++ {
        ch <- i
    }
    close(ch)
    return ch
}