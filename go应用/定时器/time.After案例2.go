package main

import (
	"fmt"
	"time"
)

func main() {
	tchan := time.After(time.Second * 3)
	fmt.Printf("tchan type=%T\n", tchan)
	fmt.Println("mark 1")
	fmt.Println("tchan=", <-tchan)
	fmt.Println("mark 2")
}
