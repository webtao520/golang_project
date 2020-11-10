排序操作和字符串格式化一样是很多程序经常使用的操作。尽管一个最短的快排程序只要 15 行就可以搞定，但是一个健壮的实现需要更多的代码，
并且我们不希望每次我们需要的时候都重写或者拷贝这些代码。

幸运的是，sort 包内置的提供了根据一些排序函数来对任何序列排序的功能。它的设计非常独到。在很多语言中，排序算法都是和序列数据类型关联，
同时排序函数和具体类型元素关联。

相比之下，Go语言的 sort.Sort 函数不会对具体的序列和它的元素做任何假设。相反，
它使用了一个接口类型 sort.Interface 来指定通用的排序算法和可能被排序到的序列类型之间的约定。
这个接口的实现由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片。

一个内置的排序算法需要知道三个东西：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式；这就是 sort.Interface 的三个方法：

package sort

type Interface interface {
    Len() int            // 获取元素数量
    Less(i, j int) bool // i，j是序列元素的指数。
    Swap(i, j int)        // 交换元素
}



为了对序列进行排序，我们需要定义一个实现了这三个方法的类型，然后对这个类型的一个实例应用 sort.Sort 函数。思考对一个字符串切片进行排序，
这可能是最简单的例子了。下面是这个新的类型 MyStringList  和它的 Len，Less 和 Swap 方法

type MyStringList  []string
func (p MyStringList ) Len() int { return len(m) }
func (p MyStringList ) Less(i, j int) bool { return m[i] < m[j] }
func (p MyStringList ) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

