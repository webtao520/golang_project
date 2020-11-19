package main

import (
    "fmt"
    "math"
)

type dualError struct {
    Num     float64
    problem string
}

func (e dualError) Error() string {
	fmt.Println("进行了自动调用了.....")
	return fmt.Sprintf("Wrong!!!,because \"%f\" is a negative number", e.Num)
	
/**
在Printf、Sprintf、Fprintf三个函数中，默认的行为是对每一个格式化verb依次对应调用时成功传递进来的参数。
但是，紧跟在verb之前的[n]符号表示应格式化第n个参数（索引从1开始）。同样的在'*'之前的[n]符号表示采用第n个参数的值作为宽度或精度。
在处理完方括号表达式[n]后，除非另有指示，会接着处理参数n+1，n+2……（就是说移动了当前处理位置）。例如：
*/
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return -1, dualError{Num: f}
    }
    return math.Sqrt(f), nil
}
func main() {
    result, err := Sqrt(-13)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }
}




/**
PS D:\goLang\github\golang_project\Go语言接口\Go语言error接口> go run 2.go
Wrong!!!,because "-13.000000" is a negative number
*/