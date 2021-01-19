package main 


import (
	"context"
	"fmt"
	 "time"
	 "go.etcd.io/etcd/clientv3"
)

// etcd  Watch

func main(){
   cli,err:=clientv3.New(clientv3.Config{
	   Endpoints:[]string{"127.0.0.1:2379"},
	   DialTimeout: 5 * time.Second,
   })
   if err !=nil {
	  // handel error 
	  fmt.Printf("connect to etcd failed, err:%v\n", err)
	  return
   }
   fmt.Println("etcd success")
   defer cli.Close()

   // watch 
   // 派一个哨兵，一直监视者  luminghui 这个key的变化，（新增，修改，删除）
   ch:=cli.Watch(context.Background(),"luminghui") // // <-chan WatchResponse  只读
   //  从通道尝试取值（监视信息）
   for wresp := range ch{
	for _, evt := range wresp.Events{
		fmt.Printf("Type:%v key:%v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
	}
}
}