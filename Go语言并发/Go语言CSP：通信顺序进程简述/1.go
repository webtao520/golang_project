package main
import (
    "fmt"
    "sync"
)
func main() {
    var mu sync.Mutex
    go func() {
        fmt.Println("C语言中文网")
        mu.Lock()
    }()
    mu.Unlock()
}