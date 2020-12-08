XML（extensible Markup Language）格式被广泛用作一种数据交换格式，并且自成一种文件格式。与上一节介绍的 JSON 相比 XML 要复杂得多，
而且手动写起来相对乏味得多。

在 JSON 还未像现在这么广泛使用时，XML 的使用相当广泛。XML 作为一种数据交换和信息传递的格式，
使用还是很广泛的，现在很多开放平台接口，基本都会支持 XML 格式。

Go语言内置的 encoding/xml 包可以用在结构体和 XML 格式之间进行编解码，其方式跟 encoding/json 包类似。
然而与 JSON 相比 XML 的编码和解码在功能上更苛刻得多，这是由于 encoding/xml 包要求结构体的字段包含格式合理的标签，而 JSON 格式却不需要。

### 写 XML 文件

使用 encoidng/xml 包可以很方便的将 xml 数据存储到文件中，示例代码如下：

+ 案例
 * 1.go

读 XML 文件
读 XML 文件比写 XML 文件稍微复杂，特别是在必须处理一些我们自定义字段的时候（例如日期）。但是，
如果我们使用合理的打上 XML 标签的结构体，就不会复杂。示例代码如下： 

+ 案例
 * 2.go

正如写 XML 时一样，我们无需关心对所读取的 XML 数据进行转义，xml.NewDecoder.Decode() 函数会自动处理这些。

xml 包还支持更为复杂的标签，包括嵌套。例如标签名为 'xml:"Books>Author"' 产生的是 <Books><Author>content</Author></Books> 这样的 XML 内容。同时除了 'xml:", attr"' 之外，该包还支持 'xml:",chardata"' 这样的标签表示将该字段当做字符数据来写，支持 'xml:",innerxml"' 这样的标签表示按照字面量来写该字段，以及 'xml:",comment"' 这样的标签表示将该字段当做 XML 注释。因此，通过使用标签化的结构体，我们可以充分利用好这些方便的编码解码函数，同时合理控制如何读写 XML 数据。 