package main 

import (
	"fmt"
	"strings"
)

func main (){
  str:="golang php"
  fmt.Printf("%v",strings.TrimLeft(str,"go"))
}