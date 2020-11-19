大家介绍一下如何把我们写好的Go语言程序发布到服务 Linux 服务器上。

下面所示的是我们在 Windows 系统下开发的代码：

###
        package main
        import (
            "fmt"
            "log"
            "net/http"
        )
        func sayHello(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintf(w, "C语言中文网")
        }
        func main() {
            http.HandleFunc("/", sayHello)
            log.Println("启动成功，可以通过 localhost:9000 访问")
            err := http.ListenAndServe(":9000", nil)
            if err != nil {
                log.Fatal("List 9000")
            }
        }
###

代码已经写好了，现在需要编译了，由于是 window 环境编译到 linux 下运行，所有涉及到跨平台编译。

编译代码命令如下所示：

###
        set GOARCH=amd64         //设置目标可执行程序操作系统构架，包括 386，amd64，arm
        set GOOS=linux           //设置可执行程序运行操作系统，支持 darwin，freebsd，linux，windows
        go build ./main.go       //打包
###

注意：使用 Window 10 系统的小伙伴必须用 cmd 工具执行上述命令，不能使用 powershell。

OK，编译完成后会生成一个 main 可执行文件，没有后缀，这时只需要把这个文件上传到你的虚拟机，直接运行就好了。

好啦！就这么简单，不需要任何语言环境，像 java 程序需要在服务器安装 java，php 需要安装 Apache，PHP 等运行环境，go 统统不需要，
只需要一个 linux 系统将编译好的代码扔上去就可以了。