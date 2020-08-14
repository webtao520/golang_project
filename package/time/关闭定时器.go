package main 

import (
	"time"
	"fmt"
)

func main(){
	ticker:=time.NewTicker(2*time.Second)
	ch:=make(chan bool)
	go func(ticker *time.Ticker){
		 defer ticker.Stop()
		 for {
			  select {
			  case <-ticker.C:
				fmt.Println("Ticker running...")
			  case stop:=<-ch:
				if stop {
					fmt.Println("Ticker Stop")
					return	
				}
			  }
		 }
	}(ticker)
	time.Sleep(10 * time.Second)
	ch<-true
	close(ch)
}

/*
Stop关闭一个Ticker。在关闭后，将不会发送更多的tick信息。Stop不会关闭通道t.C，以避免从该通道的读取不正确的成功。

Ticker running...
Ticker running...
Ticker running...
Ticker running...
Ticker running...
Ticker Stop
*/