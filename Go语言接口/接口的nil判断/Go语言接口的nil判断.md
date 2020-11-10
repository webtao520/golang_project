nil 在 Go语言中只能被赋值给指针和接口。接口在底层的实现有两个部分：type 和 data。在源码中，
显式地将 nil 赋值给接口时，
接口的 type 和 data 都将为 nil。此时，接口与 nil 值判断是相等的。但如果将一个带有类型的 nil 赋值给接口时，只有 data 为 nil，
而 type 为 nil，此时，接口与 nil 判断将不相等。

### 接口与 nil 不相等

下面代码使用 MyImplement() 实现 fmt 包中的 Stringer 接口，这个接口的定义如下：

type Stringer interface {
    String() string
}


在 GetStringer() 函数中将返回这个接口。通过 *MyImplement 指针变量置为 nil 提供 GetStringer 的返回值。在 main() 中，
判断 GetStringer 与 nil 是否相等，代码如下：
+ 案例
  * 1.go


### 为了避免这类误判的问题，可以在函数返回时，发现带有 nil 的指针时直接返回 nil，代码如下：

func GetStringer() fmt.Stringer {
    var s *MyImplement = nil
    if s == nil {
        return nil
    }
    return s
}
