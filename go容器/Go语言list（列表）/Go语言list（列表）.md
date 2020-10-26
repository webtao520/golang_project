### 初始化列表

list 的初始化有两种方法：分别是使用 New() 函数和 var 关键字声明，两种方法的初始化效果都是一致的。

1) 通过 container/list 包的 New() 函数初始化 list
变量名 := list.New()

2) 通过 var 关键字声明初始化 list
var 变量名 list.List

列表与切片和 map 不同的是，列表并没有具体元素类型的限制，因此，列表的元素可以是任意类型，这既带来了便利，也引来一些问题，例如给列表中放入了一个 interface{} 类型的值，取出值后，如果要将 interface{} 转换为其他类型将会发生宕机。
在列表中插入元素
双链表支持从队列前方或后方插入元素，分别对应的方法是 PushFront 和 PushBack。
提示
这两个方法都会返回一个 *list.Element 结构，如果在以后的使用中需要删除插入的元素，则只能通过 *list.Element 配合 Remove() 方法进行删除，这种方法可以让删除更加效率化，同时也是双链表特性之一。

下面代码展示如何给 list 添加元素：
+ 案例
    * 1.go

列表插入元素的方法如下表所示。

方  法	功  能
InsertAfter(v interface {}, mark * Element) * Element	在 mark 点之后插入元素，mark 点由其他插入函数提供

InsertBefore(v interface {}, mark * Element) *Element	在 mark 点之前插入元素，mark 点由其他插入函数提供

PushBackList(other *List)	添加 other 列表元素到尾部

PushFrontList(other *List)	添加 other 列表元素到头部    


### 从列表中删除元素

列表插入函数的返回值会提供一个 *list.Element 结构，这个结构记录着列表元素的值以及与其他节点之间的关系等信息，从列表中删除元素时，需要用到这个结构进行快速删除。

列表操作元素：
+ 案例
    * 2.go

 
列表插入元素的方法如下表所示。

方  法	功  能
InsertAfter(v interface {}, mark * Element) * Element	在 mark 点之后插入元素，mark 点由其他插入函数提供
InsertBefore(v interface {}, mark * Element) *Element	在 mark 点之前插入元素，mark 点由其他插入函数提供
PushBackList(other *List)	添加 other 列表元素到尾部
PushFrontList(other *List)	添加 other 列表元素到头部   


遍历列表——访问列表的每一个元素
遍历双链表需要配合 Front() 函数获取头元素，遍历时只要元素不为空就可以继续进行，每一次遍历都会调用元素的 Next() 函数，代码如下所示。

l := list.New()

// 尾部添加
l.PushBack("canon")

// 头部添加
l.PushFront(67)

for i := l.Front(); i != nil; i = i.Next() {
    fmt.Println(i.Value)
}

代码说明如下：
第 1 行，创建一个列表实例。
第 4 行，将 canon 放入列表尾部。
第 7 行，在队列头部放入 67。
第 9 行，使用 for 语句进行遍历，其中 i:=l.Front() 表示初始赋值，只会在一开始执行一次，每次循环会进行一次 i != nil 语句判断，如果返回 false，表示退出循环，反之则会执行 i = i.Next()。
第 10 行，使用遍历返回的 *list.Element 的 Value 成员取得放入列表时的原值。