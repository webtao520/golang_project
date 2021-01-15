package calc

import (
	"fmt"
)

func init() {
	fmt.Println("import 我时自动执行...")
}

// 包中的标识符(变量名\函数名\结构体\接口等)如果首字母是小写的,表示私有(只能在当前这个包中使用)
// 首字母大写的标识符可以被外部的包调用
func Add(x, y int) int {
	return x + y
}

func main() {

}
