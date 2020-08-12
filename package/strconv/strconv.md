### strconv 包说明

> strconv包提供了字符串与简单数据类型之间的类型转换功能。可以将简单类型转换为字符串，也可以将字符串转换为其它简单类型。

+ 这个包里提供了很多函数，大概分为几类：

    字符串转int：Atoi()
    int转字符串: Itoa()
    ParseTP类函数将string转换为TP类型：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。因为string转其它类型可能会失败，所以这些函数都有第二个返回值表示是否转换成功
    FormatTP类函数将其它类型转string：FormatBool()、FormatFloat()、FormatInt()、FormatUint()
    AppendTP类函数用于将TP转换成字符串后append到一个slice中：AppendBool()、AppendFloat()、AppendInt()、AppendUint()

+ 还有其他一些基本用不上的函数，见官方手册：go doc strconv或者https://golang.org/pkg/strconv/。

+ 当有些类型无法转换时，将报错，返回的错误是strconv包中自行定义的error类型。有两种错误：

+ var ErrRange = errors.New("value out of range")
+ var ErrSyntax = errors.New("invalid syntax")

+ 例如，使用Atoi("a")将"a"转换为int类型，自然是不成功的。如果print输出err信息，将显示：
+ strconv.Atoi: parsing "a": invalid syntax