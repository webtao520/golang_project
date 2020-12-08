JSON（JavaScript Object Notation）是一种轻量级的数据交换格式，易于阅读和编写，同时也易于机器解析和生成。
它基于 JavaScript Programming Language, Standard ECMA-262 3rd Edition - December 1999 的一个子集。

JSON 是一种使用 UTF-8 编码的纯文本格式，采用完全独立于语言的文本格式，由于写起来比 XML 格式方便，并且更为紧凑，
同时所需的处理时间也更少，致使 JSON 格式越来越流行，特别是在通过网络连接传送数据方面。

开发人员可以使用 JSON 传输简单的字符串、数字、布尔值，也可以传输一个数组或者一个更复杂的复合结构。
在 Web 开发领域中，JSON 被广泛应用于 Web 服务端程序和客户端之间的数据通信。

Go语言内建对 JSON 的支持，使用内置的 encoding/json 标准库，开发人员可以轻松使用Go程序生成和解析 JSON 格式的数据。

JSON 结构如下所示：

{"key1":"value1","key2":value2,"key3":["value3","value4","value5"]}

### 写 JSON 文件

使用Go语言创建一个 json 文件非常方便，示例代码如下：

+ 案例
  1.go


读 JSON 文件
读 JSON 数据与写 JSON 数据一样简单，示例代码如下：

+ 案例
  2.go  

顺便提一下，还有一种叫做 BSON (Binary JSON) 的格式与 JSON 非常类似，与 JSON 相比，BSON 着眼于提高存储和扫描效率。BSON 文档中的大型元素以长度字段为前缀以便于扫描。在某些情况下，由于长度前缀和显式数组索引的存在，BSON 使用的空间会多于 JSON。
