package main 

import (
	"fmt"
	"time"
)

func main () {
	now:=time.Now()
   fmt.Printf(now.Format("2006-01-02 15:04:05"))
   fmt.Println()
   fmt.Printf(now.Format("2006-01-02"))
   fmt.Println()
   fmt.Printf(now.Format("15:04:05"))
   fmt.Println()
}