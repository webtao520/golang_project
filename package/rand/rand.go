package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
}
