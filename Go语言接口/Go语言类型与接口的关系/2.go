// 使用io.Writer的代码, 并不知道Socket和io.Closer的存在
func usingWriter( writer io.Writer){
    writer.Write( nil )
}
// 使用io.Closer, 并不知道Socket和io.Writer的存在
func usingCloser( closer io.Closer) {
    closer.Close()
}
func main() {
    // 实例化Socket
    s := new(Socket)
    usingWriter(s)
    usingCloser(s)
}
// usingWriter() 和 usingCloser() 完全独立，互相不知道对方的存在，也不知道自己使用的接口是 Socket 实现的。