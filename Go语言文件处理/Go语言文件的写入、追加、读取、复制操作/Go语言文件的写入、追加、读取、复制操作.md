Go语言的 os 包下有一个 OpenFile 函数，其原型如下所示：

func OpenFile(name string, flag int, perm FileMode) (file *File, err error)

其中 name 是文件的文件名，如果不是在当前路径下运行需要加上具体路径；flag 是文件的处理参数，
为 int 类型，根据系统的不同具体值可能有所不同，但是作用是相同的。

下面列举了一些常用的 flag 文件处理参数：
O_RDONLY：只读模式打开文件；
O_WRONLY：只写模式打开文件；
O_RDWR：读写模式打开文件；
O_APPEND：写操作时将数据附加到文件尾部（追加）；
O_CREATE：如果不存在将创建一个新文件；
O_EXCL：和 O_CREATE 配合使用，文件必须不存在，否则返回一个错误；
O_SYNC：当进行一系列写操作时，每次都要等待上次的 I/O 操作完成再进行；
O_TRUNC：如果可能，在打开时清空文件。

【示例 1】：创建一个新文件 golang.txt，并在其中写入 5 句“http://c.biancheng.net/golang/”
+ 案例
  1.go

