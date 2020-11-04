package main
import (
"fmt"
"runtime"
)
func main() {
cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
fmt.Println("cpu核心数:", cpuNum)
runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量
}