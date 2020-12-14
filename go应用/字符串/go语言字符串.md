### golang语言 字符串

字符串内容不能更改，可包含byte值0，len()获取的是字节数量，
s[i]返回第i个字节的值，类型是uint8的别称byte，&s[i] 是非法的,可以作为切片使用

s:=s1+s2  会创建一个新的字符串
t:=s[3:8]  不会创建一个新的字符串，
   //应该是创建了新的直接部分，而引用相同的底层数据  待验证  源码

for i,r:=range "hello,世界" { // 遍历的是unicode 字符，不是字节
   fmt.Printf("%T",r) // int32
}   


字符串类型转换 rune是Unicode string是UTF-8

　　1，将整数值转换为字符串将产生一个包含UTF-8表示的该整数的字符串。 有效Unicode码点范围之外的值将转换为 "\uFFFD"。

　　　　string('a') // "a"

　　　　string(65) // "A"

　　　　string(-1) // "\ufffd" == "\xef\xbf\xbd "

　　　　string(0xf8) // "\u00f8" == "ø" == "\xc3\xb8"


　　2，将字节切片转换为字符串类型将产生一个连续字节为该切片元素的字符串。 若该切片值为 nil，则其结果为空字符串。

　　　　string([]byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}) // "hellø"

　　　　string([]byte{ '啊' }) // 报错 constant 21834 overflows byte

　　　　string([]byte{}) // ""

　　　　string([]byte(nil)) // ""

　　　　type MyBytes []byte

　　　　string(MyBytes{'h', 'e', 'l', 'l', '\xc3', '\xb8'}) // "hellø"


　　3，将字符串类型值转换为字节类型切片将产生一个连续元素为该字符串字节的切片

　　　　[]byte("hellø") // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}

　　　　[]byte("") // []byte{}

　　　　MyBytes("hellø") //[]byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}


　　4，将符文切片转换为字符串类型，将对rune进行utf8编码。 若该切片值为 nil，则其结果为空字符串。

　　　　string([]rune{0x767d, 0x9d6c, 0x7fd4}) // "\u767d\u9d6c\u7fd4" == "白鵬翔"

　　　　string([]rune{}) // ""

　　　　string([]rune(nil)) // ""

　　　　type MyRunes []rune

　　　　string(MyRunes{0x767d, 0x9d6c, 0x7fd4}) // "\u767d\u9d6c\u7fd4" == "白鵬翔"


　　5，将字符串类型的值转换为符文类型切片将产生一个包含该字符串单个Unicode码点的切片

　　　　[]rune(MyString("白鵬翔")) // []rune{0x767d, 0x9d6c, 0x7fd4}

　　　　[]rune("") // []rune{}

　　　　MyRunes("白鵬翔") // []rune{0x767d, 0x9d6c, 0x7fd4}


### Rune字面量

由包含在单引号中的若干个字符表示，除了换行符和未转义的单引号

引号内的单个字符表示该字符本身的Unicode值，如 'a'，'啊'

以反斜杠开头的多个字符以多种形式编码为值，如 ‘\n’，'\377'，'\x07'，'\u12e4'，'\U00101234'

小于256的码点值可以用16进制转义字节中'\xab'，更大的码点必须用\u形式，因此\xe4\xb8\x96不是一个合法的rune字符，尽管这三个字节对应一个有效的utf8编码的码点。

rune类型的零值常用 '\000'、'\x00'或'\u0000'等来表示

'a' == 97 //true

 

### 字符串字面量

原始字符串(raw string)字面量是反引号之间的字符序列，在该引号内， 除反引号外的任何字符都是合法的。反斜杠没有特殊意义，回车符\r会被从值中丢弃。

解译字符串(Interpreted string)字面量是双引号之间的字符序列，双引号内不能有换行符和未转义的双引号。反斜杠的作用类似Rune。“\xe4\xb8\x96”表示字符串"世"

字符串零值为""或``
