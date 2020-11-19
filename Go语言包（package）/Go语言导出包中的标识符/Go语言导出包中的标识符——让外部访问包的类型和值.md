在 Go语言中，如果想在一个包里引用另外一个包里的标识符（如类型、变量、常量等）时，必须首先将被引用的标识符导出，
将要导出的标识符的首字母大写就可以让引用者可以访问这些标识符了。

### 导出包内标识符

下面代码中包含一系列未导出标识符，它们的首字母都为小写，这些标识符可以在包内自由使用，但是包外无法访问它们，代码如下：
###
       package mypkg 
       var myVar = 100
       const myConst = "hello"
       type myStruct struct  {

       }
###
将 myStruct 和 myConst 首字母大写，导出这些标识符，修改后代码如下：

###
    package mypkg 
    var myVar = 100
    const MyConst  = "hello"
    type MyStruct struct  {

       }
###

此时，MyConst 和 MyStruct 可以被外部访问，而 myVar 由于首字母是小写，因此只能在 mypkg 包内使用，不能被外部包引用。

### 导出结构体及接口成员
在被导出的结构体或接口中，如果它们的字段或方法首字母是大写，外部可以访问这些字段和方法，代码如下：
### 
        type MyStruct struct {
            // 包外可以访问的字段
            ExportedField int
            // 仅限包内访问的字段
            privateField int
        }
        type MyInterface interface {
            // 包外可以访问的方法
            ExportedMethod()
            // 仅限包内访问的方法
            privateMethod()
        }
### 