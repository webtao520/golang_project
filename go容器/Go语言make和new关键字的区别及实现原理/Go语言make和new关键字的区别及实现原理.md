Go语言中 new 和 make 是两个内置函数，主要用来创建并分配类型的内存。在我们定义变量的时候，
可能会觉得有点迷惑，不知道应该使用哪个函数来声明变量，其实他们的规则很简单，new 只分配内存，
而 make 只能用于 slice、map 和 channel 的初始化，下面我们就来具体介绍一下。

### new

在Go语言中，new 函数描述如下：

// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
func new(Type) *Type

从上面的代码可以看出，new 函数只接受一个参数，这个参数是一个类型，并且返回一个指向该类型内存地址的指针。同时 new 函数会把分配的内存置为零，也就是类型的零值。

【示例】使用 new 函数为变量分配内存空间。
+ 案例
    * 1.go

当然，new 函数不仅仅能够为系统默认的数据类型，分配空间，自定义类型也可以使用 new 函数来分配空间，如下所示：
+ 案例
    * 2.go

这就是 new 函数，它返回的永远是类型的指针，指针指向分配类型的内存地址。

### make

make 也是用于内存分配的，但是和 new 不同，它只用于 chan、map 以及 slice 的内存创建，
而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

在Go语言中，make 函数的描述如下：

// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not a
// value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it. The specification of the result depends on
// the type:
// Slice: The size specifies the length. The capacity of the slice is
// equal to its length. A second integer argument may be provided to
// specify a different capacity; it must be no smaller than the
// length, so make([]int, 0, 10) allocates a slice of length 0 and
// capacity 10.
// Map: An empty map is allocated with enough space to hold the
// specified number of elements. The size may be omitted, in which case
// a small starting size is allocated.
// Channel: The channel's buffer is initialized with the specified
// buffer capacity. If zero, or the size is omitted, the channel is
// unbuffered.
func make(t Type, size ...IntegerType) Type
通过上面的代码可以看出 make 函数的 t 参数必须是 chan（通道）、map（字典）、slice（切片）中的一个，并且返回值也是类型本身。
注意：make 函数只用于 map，slice 和 channel，并且不返回指针。如果想要获得一个显式的指针，可以使用 new 函数进行分配，或者显式地使用一个变量的地址。

Go语言中的 new 和 make 主要区别如下：
make 只能用来分配及初始化类型为 slice、map、chan 的数据。new 可以分配任意类型的数据；

new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type；

new 分配的空间被清零。make 分配空间后，会进行初始化；













