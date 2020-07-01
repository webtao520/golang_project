## 先举个例子：

> 在 Go http包的Server中，每一个请求在都有一个对应的 goroutine 去处理。请求处理函数通常会启动额外的 goroutine 用来访问后端服务，比如数据库和RPC服务。用来处理一个请求的 goroutine 通常需要访问一些与请求特定的数据，比如终端用户的身份认证信息、验证相关的token、请求的截止时间。 当一个请求被取消或超时时，所有用来处理该请求的 goroutine 都应该迅速退出，然后系统才能释放这些 goroutine 占用的资源。

对于多并发的情况下，传统的方案：等待组sync.WaitGroup以及通过通道channel的方式的问题就会显现出来；

对于等待组控制多并发的情况：只有所有的goroutine都结束了才算结束，只要有一个goroutine没有结束， 那么就会一直等，这显然对资源的释放是缓慢的；

而对于通道Channel的方式下：通过在main goroutine中像chan中发送关闭停止指令，并配合select，从而达到关闭goroutine的目的，这种方式显然比等待组优雅的多，但是在goroutine中在嵌套goroutine的情况就变得异常复杂。

##### 等待组例子：
 > 见 context.go 

 上面例子中，一个任务结束了必须等待另外一个任务也结束了才算全部结束了，先完成的必须等待其他未完成的，所有的goroutine都要全部完成才OK。

这种方式的优点：使用等待组的并发控制模型，尤其适用于好多个goroutine协同做一件事情的时候，因为每个goroutine做的都是这件事情的一部分，只有全部的goroutine都完成，这件事情才算完成；

这种方式的缺陷：在实际生产中，需要我们主动的通知某一个 goroutine 结束。

> 比如我们开启一个后台 goroutine 一直做事情，比如监控，现在不需要了，就需要通知这个监 goroutine 结束，不然它会一直跑，就泄漏了。

## 通道+select的方式

  + 在等待组例子的最后抛出了一个问题，针对这种问题有2种办法：

    * 设置全局变量，在我们需要通知goroutine要停止的时候，我们为全局变量赋值，但是这样我们必须保证线程安全，不可避免的我们要为全局变量加锁，在便利性及性能上稍显不足；

    * 使用chan+select多路复用的方式，就会优雅许多；

##### chan+select 案例见 chan_select.go

上面例子中：我们定义一个 stop 的 chan，通知它结束后台 goroutine。

实现也非常简单，在后台 goroutine 中，使用 select 判断 stop 是否可以接收到值，如果可以接收到，就表示可以退出停止了；如果没有接收到，就会执行 default 里的逻辑，继续运行，直到收到 stop 的通知。

发送了 stop<- true 结束的指令后，我这里使用 time.Sleep(3 * time.Second) 故意停顿 3 秒来检测我们结束任务goroutine 是否成功。

如果成功的话，不会再有 "任务1 正在运行中。" 的输出了；如果没有成功，监控 goroutine 就会继续打印 "任务1 正在运行中。" 输出。

channel配合select方式的优点：比较优雅，

channel配合select方式的劣势：如果有很多 goroutine 都需要控制结束怎么办？， 如果这些 goroutine 又衍生了其它更多的goroutine 怎么办？


## context的加入

context是GO1.7版本加入的一个标准库，它定义了Context类型，专门用来简化 对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。

对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。它们之间的函数调用链必须传递上下文，或者可以使用WithCancel、WithDeadline、WithTimeout或WithValue创建的派生上下文。当一个上下文被取消时，它派生的所有上下文也被取消。

当一个goroutine在衍生一个goroutine时，context可以跟踪到子goroutine，从而达到控制他们的目的；

##### 使用context重写上面的select例子 见 context_select.go

重写比较简单，就是把原来的 chan stop 换成 Context，使用 Context 跟踪 goroutine，以便进行控制，比如结束等。

context.Background() 返回一个空的 Context，这个空的 Context 一般用于整个 Context 树的根节点。然后我们使用 context.WithCancel(parent) 函数，创建一个可取消的子 Context，然后当作参数传给 goroutine 使用，这样就可以使用这个子 Context 跟踪这个 goroutine。

在 goroutine 中，使用 select 调用<-ctx.Done()判断是否要结束，如果接受到值的话，就可以返回结束 goroutine 了；如果接收不到，就会继续进行运行任务。

那么是如何发送结束指令的呢？这就是示例中的 cancel 函数啦，它是我们调用context.WithCancel(parent) 函数生成子 Context 的时候返回的，第二个返回值就是这个取消函数，它是 CancelFunc 类型的。我们调用它就可以发出取消指令，然后我们的监控 goroutine 就会收到信号，就会返回结束。


##### Context控制多个goroutine  见 context_all_goroutine.go

上面例子中，启动了 3 个监控 goroutine 进行不断的运行任务，每一个都使用了 Context 进行跟踪，当我们使用 cancel 函数通知取消时，这 3 个 goroutine 都会被结束。这就是 Context 的控制能力，它就像一个控制器一样，按下开关后，所有基于这个 Context 或者衍生的子 Context 都会收到通知，这时就可以进行清理操作了，最终释放 goroutine，这就优雅的解决了 goroutine 启动后不可控的问题。


## context接口

          type Context interface {
              Deadline() (deadline time.Time, ok bool)
              Done() <- chan struct {}
              Err() error 
              Value(key interface{}) interface{}
          }

+ Context 的接口定义的比较简洁，这个接口共有 4 个方法;

 * Deadline:是获取设置的截止时间的意思，第一个返回值是截止时间，到了这个时间点，Context 会自动发起取消请求；第二个返回值 ok==false 时表示没有设置截止时间，如果需要取消的话，需要调用取消函数进行取消。

 * Done:该方法返回一个只读的 chan，类型为 struct{}，我们在 goroutine 中，如果该方法返回的 chan 可以读取，则意味着parent context已经发起了取消请求，我们通过 Done 方法收到这个信号后，就应该做清理操作，然后退出 goroutine，释放资源。

 * Err 方法返回取消的错误原因，因为什么 Context 被取消。

 * Value方法获取该 Context 上绑定的值，是一个键值对，所以要通过一个 Key 才可以获取对应的值，这个值一般是线程安全的。 

 四个方法中常用的就是 Done 了，如果 Context 取消的时候，我们就可以得到一个关闭的 chan，关闭的 chan 是可以读取的，所以只要可以读取的时候，就意味着收到 Context 取消的信号了，以下是这个方法的经典用法。

        func  Stream (ctx context.Context, out chan <- Value)error {
            for {
                v,err:=DoSomethine(ctx)
                if err !=nil {
                    return err
                }
                select {
                    case <-ctx.Done():
                    return ctx.Err()
                    case out <- v:
                }
            }
        }

Context 接口并不需要我们实现，Go 内置已经帮我们实现了 2 个，我们代码中最开始都是以这两个内置的作为最顶层的 partent context，衍生出更多的子 Context。

        var (
            background = new(emptyCtx)
            todo       = new(emptyCtx)
        )

        func Background() Context {
            return background
        }

        func TODO() Context {
            return todo
        }

+ 
    * Background()主要用于 main 函数、初始化以及测试代码中，作为 Context 这个树结构的最顶层的 Context，也就是根 Context。

    * TODO()，它目前还不知道具体的使用场景，如果我们不知道该使用什么 Context 的时候，可以使用这个。

它们两个本质上都是 emptyCtx 结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的 Context。

        type emptyCtx int

        func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
            return
        }

        func (*emptyCtx) Done() <-chan struct{} {
            return nil
        }

        func (*emptyCtx) Err() error {
            return nil
        }

        func (*emptyCtx) Value(key interface{}) interface{} {
            return nil
        }


这就是 emptyCtx 实现 Context 接口的方法，可以看到，这些方法什么都没做，返回的都是 nil 或者零值。


## Context的继承衍生

有了如上的根 Context，那么是如何衍生更多的子 Context 的呢？这就要靠 context 包为我们提供的 With 系列的函数了。

  func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
  func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
  func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
  func WithValue(parent Context, key, val interface{}) Context

+ 这四个 With 函数，接收的都有一个 partent 参数，就是父 Context，我们要基于这个父 Context 创建出子 Context 的意思，这种方式可以理解为子 Context 对父 Context 的继承，也可以理解为基于父 Context 的衍生。通过这些函数，就创建了一颗 Context 树，树的每个节点都可以有任意多个子节点，节点层级可以有任意多个。

* WithCancel 函数，传递一个父 Context 作为参数，返回子 Context，以及一个取消函数用来取消 Context。 WithDeadline 函数，和 WithCancel 差不多，它会多传递一个截止时间参数，意味着到了这个时间点，会自动取消 Context，当然我们也可以不等到这个时候，可以提前通过取消函数进行取消。

* WithTimeout 和WithDeadline 基本上一样，这个表示是超时自动取消，是多少时间后自动取消 Context 的意思。

* WithValue 函数和取消 Context 无关，它是为了生成一个绑定了一个键值对数据的 Context，即给context设置值，这个绑定的数据可以通过 Context.Value 方法访问到.

上面3个函数都会返回一个取消函数CancelFunc，这是一个函数类型，它的定义非常简单type CancelFunc func(),该函数可以取消一个 Context，以及这个节点 Context下所有的所有的 Context，不管有多少层级。

##### Context传递元数据  见 context_val.go

+ 注意点

* context.WithValue 方法附加一对 K-V 的键值对，这里 Key 必须是等价性的，也就是具有可比性；Value 值要是线程安全的。

* 在使用值的时候，可以通过 Value 方法读取 ctx.Value(key)。

* 使用 WithValue 传值，一般是必须的值，不要什么值都传递。

+ Context最佳实战

* 不要把 Context 放在结构体中，要以参数的方式传递

* 以 Context 作为参数的函数方法，应该把 Context 作为第一个参数，放在第一位

* 给一个函数方法传递 Context 的时候，不要传递 nil，如果不知道传递什么，就使用 context.TODO

* Context 的 Value 相关方法应该传递必须的数据，不要什么数据都使用这个传递

* Context 是线程安全的，可以放心的在多个 goroutine 中传递


# 参考博客：
  + https://zhuanlan.zhihu.com/p/58967892
  
  + https://www.liwenzhou.com/posts/Go/go_context/