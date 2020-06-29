/*
	# Go 语言错误处理机制 #

	Go 语言为错误处理定义了一个标准模式，即 error 接口，该接口的定义非常简单：

	type error interface {
		Error() string
	}

	其中只包含了一个 Error() 方法，用于返回错误消息。对于大多数函数或类方法，如果要返回错误，
	基本都可以定义成如下模式 —— 将错误类型作为第二个参数返回：

	func Foo(param int)(n int,err error){
		//....
	}

	然后在调用返回错误信息的函数/方法时，按照如下「卫述语句」模板编写处理代码即可：

	 n,err:= Foo(0)
	 if err!=nil {
		 // 错误处理
	 }else{
		 // 使用返回值 n
	 }

    非常简洁优雅

	# 错误消息返回及处理 #

	关于自定义并返回 error 类型错误信息的使用示例，前面介绍函数返回值时已经演示过，
	我们可以通过 errors.New() 方法创建 error 类型错误信息：

	func  add(a,b *int)(c int,err error){
		if (*a < 0 || *b < 0){
			err= errors.New("只支持非负整数相加")
			return
		}
		*a *=2
		*b *=3
		c=*a+*b
		return
	}

	调用这个函数并处理错误的代码如下所示：

		x, y := 1, 2
		z, err := add(&x, &y)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("add(%d, %d) = %d\n", x, y, z)
		}

		注意到我们在打印错误信息时，直接传入了 err 对象实例，因为 Go 底层会自动调用 err 实例上的 Error() 方法返回错误信息并将其打印出来，
		就像普通类的 String() 方法一样。

		以上这种错误处理已经能够满足我们日常编写 Go 代码时大部分错误处理的需求了，事实上，
		Go 底层很多包进行错误处理时就是这样做的，此外我们还可以通过 fmt.Errorf() 格式化方法返回 error 类型错误，
		其底层调用的其实也是 errors.New 方法：

		func Errorf(format string, a ...interface{}) error {
    		return errors.New(Sprintf(format, a...))
		}

		# 内置错误类型 #

		除了上面这种最基本的使用 errors.New 方法返回错误信息之外，Go 语言内置的很多包还封装了很多更复杂的错误类型，
		以 os 包为例，这个包主要与操作平台的文件系统打交道，所以提供了 LinkError、PathError、
		SyscallError 这些实现了 error 接口的错误类型，以 PathError 为例，其底层类型结构如下：

		type PathError struct {
			Op   string
			Path string
			Err  error
		}

		该错误类型除了组合 error 接口实现 Error() 方法外，还提供了额外的操作类型字段 Op 和文件路径字段 Path 以丰富错误信息，
		该类型的 Error() 方法实现如下：

		func (e *PathError) Error() string {
			return e.Op + " " + e.Path + ": " + e.Err.Error()
		}

		我们可以在调用 os 包方法出错时通过 switch 分支语句判定具体的错误类型进行处理：

		fi, err := os.Stat("test.txt")
			if err != nil {
				switch err := err.(type) {
				case *os.PathError:
					// do something
				case *os.LinkError:
					// dome something
				case *os.SyscallError:
					// dome something
				case *exec.Error:
					// dome something
				}
			} else {
				// ...
			}

			# 自定义错误类型 #
				当然，我们也可以仿照 PathError 的实现自定义一些复杂的错误类型，只需要组合 error 接口实现 Error() 方法即可，
				然后按照自己的需要为自定义类型添加一些属性字段，很简单，这里就不展开介绍了。

*/
package main

func main() {

}
