package main 
import (
	"fmt"
	"time"
)


// Duration类型代表两个时间点之间经过的时间，以纳秒为单位。可表示的最长时间段大约290年
func main () {
	//要将Duration类型值表示为某时间单元的个数，用除法：

	second := time.Second
	fmt.Print(int64(second/time.Millisecond)) // prints 1000

	//要将整数个某时间单元表示为Duration类型值，用乘法：
	fmt.Println("\n")
	seconds := 10
	fmt.Print(time.Duration(seconds)*time.Second) // prints 10s
}