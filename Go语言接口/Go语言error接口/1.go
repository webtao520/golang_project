package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, errors.New("math: square root of negative number")
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
PS D:\goLang\github\golang_project\Go语言接口\Go语言error接口> go run 1.go
math: square root of negative number
*/
