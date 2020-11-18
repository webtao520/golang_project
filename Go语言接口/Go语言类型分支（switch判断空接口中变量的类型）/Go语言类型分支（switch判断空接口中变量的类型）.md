type-switch 流程控制的语法或许是Go语言中最古怪的语法。 它可以被看作是类型断言的增强版。
它和 switch-case 流程控制代码块有些相似。 一个 type-switch 流程控制代码块的语法如下所示：

switch t := areaIntf.(type) {
case *Square:
    fmt.Printf("Type Square %T with value %v\n", t, t)
case *Circle:
    fmt.Printf("Type Circle %T with value %v\n", t, t)
case nil:
    fmt.Printf("nil value: nothing to check?\n")
default:
    fmt.Printf("Unexpected type %T\n", t)
}


输出结构如下：
Type Square *main.Square with value &{5}

变量 t 得到了 areaIntf 的值和类型， 所有 case 语句中列举的类型（nil 除外）都必须实现对应的接口，
如果被检测类型没有在 case 语句列举的类型中，就会执行 default 语句。

如果跟随在某个 case 关键字后的条目为一个非接口类型（用一个类型名或类型字面表示），
则此非接口类型必须实现了断言值 x 的（接口）类型。


### 类型断言的书写格式

switch 实现类型分支时的写法格式如下：

switch 接口变量.(type) {
    case 类型1:
        // 变量是类型1时的处理
    case 类型2:
        // 变量是类型2时的处理
    …
    default:
        // 变量不是所有case中列举的类型时的处理
}
对各个部分的说明：
接口变量：表示需要判断的接口类型的变量。
类型1、类型2……：表示接口变量可能具有的类型列表，满足时，会指定 case 对应的分支进行处理。


### 使用类型分支判断基本类型

下面的例子将一个 interface{} 类型的参数传给 printType() 函数，通过 switch 判断 v 的类型，然后打印对应类型的提示，代码如下：
+ 案例
 * 1.go

代码第 9 行中，v.(type) 就是类型分支的典型写法。通过这个写法，在 switch 的每个 case 中写的将是各种类型分支。

代码经过 switch 时，会判断 v 这个 interface{} 的具体类型从而进行类型分支跳转。

switch 的 default 也是可以使用的，功能和其他的 switch 一致。


使用类型分支判断接口类型
多个接口进行类型断言时，可以使用类型分支简化判断过程。

现在电子支付逐渐成为人们普遍使用的支付方式，电子支付相比现金支付具备很多优点。
例如，电子支付能够刷脸支付，而现金支付容易被偷等。使用类型分支可以方便地判断一种支付方法具备哪些特性，具体请参考下面的代码。

+ 案例
 * 2.go



