package main
import (
    "fmt"
)
func main() {
    var m map[int]string
    var ptr *int
    fmt.Printf(m == ptr)
}

/**
PS D:\goLang\github\golang_project\go容器\Go语言nil：空值\零值\Go语言nil：空值> go run 4.go
# command-line-arguments
.\4.go:8:18: invalid operation: m == ptr (mismatched types map[int]string and *int)
*/