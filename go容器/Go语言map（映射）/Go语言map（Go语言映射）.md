Go语言中 map 是一种特殊的数据结构，一种元素对（pair）的无序集合，pair 对应一个 key（索引）和一个 value（值），
所以这个结构也称为关联数组或字典，这是一种能够快速寻找值的理想结构，给定 key，就可以迅速找到对应的 value。

map 这种数据结构在其他编程语言中也称为字典（Python）、hash 和 HashTable 等。

### map 概念

map 是引用类型，可以使用如下方式声明：
var mapname map[keytype]valuetype

其中：
mapname 为 map 的变量名。
keytype 为键类型。
valuetype 是键对应的值类型。
提示：[keytype] 和 valuetype 之间允许有空格。

在声明的时候不需要知道 map 的长度，因为 map 是可以动态增长的，未初始化的 map 的值是 nil，使用函数 len() 可以获取 map 中 pair 的数目。

+ 案例
    * 1.go

示例中 mapLit 演示了使用{key1: value1, key2: value2}的格式来初始化 map ，就像数组和结构体一样。

上面代码中的 mapCreated 的创建方式
mapCreated := make(map[string]float)等价于mapCreated := map[string]float{} 。


mapAssigned 是 mapList 的引用，对 mapAssigned 的修改也会影响到 mapLit 的值。

注意：可以使用 make()，但不能使用 new() 来构造 map，如果错误的使用 new() 分配了一个引用对象，会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：
mapCreated := new(map[string]float)

接下来当我们调用mapCreated["key1"] = 4.5的时候，编译器会报错：
invalid operation: mapCreated["key1"] (index of type *map[string]float).    


### map 容量

和数组不同，map 可以根据新增的 key-value 动态的伸缩，因此它不存在固定长度或者最大限制，但是也可以选择标明 map 的初始容量 capacity，格式如下：
make(map[keytype]valuetype, cap)

例如：
map2 := make(map[string]float, 100)

当 map 增长到容量上限的时候，如果再增加新的 key-value，map 的大小会自动加 1，所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。

这里有一个 map 的具体例子，即将音阶和对应的音频映射起来：

+ 案例
    * 2.go

### 用切片作为 map 的值

既然一个 key 只能对应一个 value，而 value 又是一个原始类型，那么如果一个 key 要对应多个值怎么办？例如，
当我们要处理 unix 机器上的所有进程，以父进程（pid 为整形）作为 key，
所有的子进程（以所有子进程的 pid 组成的切片）作为 value。通过将 value 定义为 []int 类型或者其他类型的切片，
就可以优雅的解决这个问题，示例代码如下所示：   

mp1 := make(map[int][]int)
mp2 := make(map[int]*[]int)



