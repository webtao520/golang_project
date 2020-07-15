# go context 讲解

> 控制并发有两种经典的方式，一种是WaitGroup，另外一种就是Context，今天我就谈谈Context。

+ 什么是WaitGroup

  * WaitGroup它是一种控制并发的方式，它的这种方式是控制多个goroutine同时完成。
  * 这是一种控制并发的方式，这种尤其适用于，好多个goroutine协同做一件事情的时候，因为每个goroutine做的都是这件事情的一部分，只有全部的goroutine都完成，这件事情才算是完成，这是等待的方式。一个很简单的例子，一定要例子中的2个goroutine同时做完，才算是完成，先做好的就要等着其他未完成的，所有的goroutine要都全部完成才可以。

  > 案例代码 WaitGroup/WaitGroup.go
      
> 在实际的业务种，我们可能会有这么一种场景：需要我们主动的通知某一个goroutine结束。比如我们开启一个后台goroutine一直做事情，比如监控，现在不需要了，就需要通知这个监控goroutine结束，不然它会一直跑，就泄漏了。


+ chan通知

 * 我们都知道一个goroutine启动后，我们是无法控制他的，大部分情况是等待它自己结束，那么如果这个goroutine是一个不会自己结束的后台goroutine呢？比如监控等，会一直运行的。

 * 这种情况化，一直傻瓜式的办法是全局变量，其他地方通过修改这个变量完成结束通知，然后后台goroutine不停的检查这个变量，如果发现被通知关闭了，就自我结束。

 * 这种方式也可以，但是首先我们要保证这个变量在多线程下的安全，基于此，有一种更好的方式：chan + select 。

   > 案例代码 WaitGroup/chan+select.go

   > 有了以上的逻辑，我们就可以在其他goroutine种，给stop chan发送值了，例子中是在main goroutine中发送的，控制让这个监控的goroutine结束。例子中我们定义一个stop的chan，通知他结束后台goroutine。实现也非常简单，在后台goroutine中，使用select判断stop是否可以接收到值，如果可以接收到，就表示可以退出停止了；如果没有接收到，就会执行default里的监控逻辑，继续监控，只到收到stop的通知。发送了stop<- true结束的指令后，我这里使用time.Sleep(5 * time.Second)故意停顿5秒来检测我们结束监控goroutine是否成功。如果成功的话，不会再有goroutine监控中...的输出了；如果没有成功，监控goroutine就会继续打印goroutine监控中...输出。这种chan+select的方式，是比较优雅的结束一个goroutine的方式，不过这种方式也有局限性，如果有很多goroutine都需要控制结束怎么办呢？如果这些goroutine又衍生了其他更多的goroutine怎么办呢？如果一层层的无穷尽的goroutine呢？这就非常复杂了，即使我们定义很多chan也很难解决这个问题，因为goroutine的关系链就导致了这种场景非常复杂。


+ 初识Context

  *上面说的这种场景是存在的，比如一个网络请求Request，每个Request都需要开启一个goroutine做一些事情，这些goroutine又可能会开启其他的goroutine。所以我们需要一种可以跟踪goroutine的方案，才可以达到控制他们的目的，这就是Go语言为我们提供的Context，称之为上下文非常贴切，它就是goroutine的上下文。
  
  * 下面我们就使用Go Context重写上面的示例。

    > 案例代码 Context/context.go

  > context.Background() 返回一个空的Context，这个空的Context一般用于整个Context树的根节点。然后我们使用context.WithCancel(parent)函数，创建一个可取消的子Context，然后当作参数传给goroutine使用，这样就可以使用这个子Context跟踪这个goroutine。重写比较简单，就是把原来的chan stop 换成Context，使用Context跟踪goroutine，以便进行控制，比如结束等。在goroutine中，使用select调用<-ctx.Done()判断是否要结束，如果接受到值的话，就可以返回结束goroutine了；如果接收不到，就会继续进行监控。那么是如何发送结束指令的呢？这就是示例中的cancel函数啦，它是我们调用context.WithCancel(parent)函数生成子Context的时候返回的，第二个返回值就是这个取消函数，它是CancelFunc类型的。我们调用它就可以发出取消指令，然后我们的监控goroutine就会收到信号，就会返回结束。

+ Context控制多个goroutine

* 使用Context控制一个goroutine的例子如上，非常简单，下面我们看看控制多个goroutine的例子，其实也比较简单。
  
   > 案例代码 Context/contexts.go

