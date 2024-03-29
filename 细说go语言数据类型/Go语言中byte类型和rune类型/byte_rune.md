+ go语言中字符串需要使用用双引号，而单引号用来表示单个的字符，字符也是组成字符串的元素。go语言的字符有两种：

    * uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
    * rune类型，代表一个 UTF-8字符。
    * ASCII编码是1个字节，而UTF-8是可变长的编码，当要表示中文等非ASCll编码的字符时，需要使用UTF-8编码来保证不会乱码

+ 举个例子，来遍历打印一个字符串，当使用byte类型时： 见byte_rune.go

    * 上面结果，英文字符正确打印，但中文乱码。是因为UTF8编码下一个中文汉字由3~4个字节组成，而字符串是由byte字节组成，所以长度也是byte字符长度，这样遍历时遇到中文就乱码了。


    * golang中string底层是通过byte数组实现的。中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。

+ 遇到这种带中文的字符串，可以使用go提供的另一个方法来遍历：

    * 使用range，其实是使用rune类型来编码的，rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。


#### 修改字符串

+ 所谓对字符串的修改其实不是对字符串本身的修改，而是复制字符串，同时修改值，即重新分配来内存。

+ 在go中修改字符串，需要先将字符串转化成数组，[]byte 或 []rune，然后再转换成 string型。
    
    * 对于全是ASCII编码的字符串： 见 ascii.go


#### rune数据类型
+ 查询,官方的解释如下：
    // rune is an alias for int32 and is equivalent to int32 in all ways. It is
    // used, by convention, to distinguish character values from integer values.

    //int32的别名，几乎在所有方面等同于int32
    //它用来区分字符值和整数值

    type rune = int32

    案例见 rune.go

> golang中string底层是通过byte数组实现的。中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。

> 我们预期想得到一个字符串的长度，而不是字符串底层占得字节长度，该怎么办呢？？？

   案例见附件 len_rune_byte.go
   

