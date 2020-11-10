package  main 

import (
	"fmt"
)

type Socket  struct {

}

func (s *Socket) Write (p []byte) (n int,err error){
	return 0,nil
}

func (s *Socket) Close() error {
	return nil
}

func main(){
   fmt.println("一个类型可以实现多个接口")
}