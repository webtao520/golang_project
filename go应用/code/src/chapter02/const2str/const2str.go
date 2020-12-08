package main

import "fmt"

// 声明芯片类型
type ChipType int

const (
	None ChipType = iota
	CPU           // 中央处理器
	GPU           // 图形处理器
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}

	return "N/A"
}

func main() {

	// 输出CPU的值和以整形方式显示
	fmt.Printf("%s %d", CPU, CPU)
}
