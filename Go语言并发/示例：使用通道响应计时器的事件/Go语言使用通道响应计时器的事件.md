Go语言中的 time 包提供了计时器的封装。由于 Go语言中的通道和 goroutine 的设计，定时任务可以在 goroutine 中通过同步的方式完成，
也可以通过在 goroutine 中异步回调完成。这里将分两种用法进行例子展示。

一段时间之后（time.After）
延迟回调：
 + 案例
    * 1.go 

time.AfterFunc() 函数是在 time.After 基础上增加了到时的回调，方便使用。

而 time.After() 函数又是在 time.NewTimer() 函数上进行的封装，下面的例子展示如何使用 timer.NewTimer() 和 time.NewTicker()。    


### 定点计时

计时器（Timer）的原理和倒计时闹钟类似，都是给定多少时间后触发。打点器（Ticker）的原理和钟表类似，钟表每到整点就会触发。
这两种方法创建后会返回 time.Ticker 对象和 time.Timer 对象，里面通过一个 C 成员，类型是只能接收的时间通道（<-chan Time），
使用这个通道就可以获得时间触发的通知。

下面代码创建一个打点器，每 500 毫秒触发一起；创建一个计时器，2 秒后触发，只触发一次。
计时器：
 + 案例
    * 2.go 