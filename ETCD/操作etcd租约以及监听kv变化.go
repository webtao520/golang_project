package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.etcd.io/etcd/clientv3"
)

const (
	NewLeaseErr  = 101
	LeasTtlErr   = 102
	KeepAliveErr = 103
	PutErr       = 104
	GetErr       = 105
	RevokeErr    = 106
)

func main() {
	var conf = clientv3.Config{
		Endpoints:   []string{"172.16.196.129:2380", "192.168.50.250:2380"},
		DialTimeout: 5 * time.Second,
	}
	client, err := clientv3.New(conf)
	defer client.Close()
	if err != nil {
		fmt.Printf("创建client失败:\n", err.Error())
		os.Exit(NewLeaseErr)
	}
	//创建租约
	lease := clientv3.NewLease(client)
	//设置租约时间
	leaseResp, err := lease.Grant(context.TODO(), 10)
	if err != nil {
		fmt.Printf("设置租约时间失败:%s\n", err.Error())
		os.Exit(LeasTtlErr)
	}
	//设置续租
	leaseID := leaseResp.ID
	ctx, cancelFunc := context.WithCancel(context.TODO())
	leaseRespChan, err := lease.KeepAlive(ctx, leaseID)
	if err != nil {
		fmt.Printf("续租失败:%s\n", err.Error())
		os.Exit(KeepAliveErr)
	}
	//监听租约
	go func() {
		for {
			select {
			case leaseKeepResp := <-leaseRespChan:
				if leaseKeepResp == nil {
					fmt.Printf("已经关闭续租功能\n")
					return
				} else {
					fmt.Printf("续租成功\n")
					goto END
				}
			}
		END:
			time.Sleep(500 * time.Millisecond)
		}
	}()
	//监听某个key的变化
	//ctx1, _ := context.WithTimeout(context.TODO(),20)
	go func() {
		wc := client.Watch(context.TODO(), "/job/v3/1", clientv3.WithPrevKV())
		for v := range wc {
			for _, e := range v.Events {
				fmt.Printf("type:%v kv:%v  prevKey:%v \n ", e.Type, string(e.Kv.Key), e.PrevKv)
			}
		}
	}()
	kv := clientv3.NewKV(client)
	//通过租约put
	putResp, err := kv.Put(context.TODO(), "/job/v3/1", "koock", clientv3.WithLease(leaseID))
	if err != nil {
		fmt.Printf("put 失败：%s", err.Error())
		os.Exit(PutErr)
	}
	fmt.Printf("%v\n", putResp.Header)
	cancelFunc()
	time.Sleep(2 * time.Second)
	_, err = lease.Revoke(context.TODO(), leaseID)
	if err != nil {
		fmt.Printf("撤销租约失败:%s\n", err.Error())
		os.Exit(RevokeErr)
	}
	fmt.Printf("撤销租约成功")
	getResp, err := kv.Get(context.TODO(), "/job/v3/1")
	if err != nil {
		fmt.Printf("get 失败：%s", err.Error())
		os.Exit(GetErr)
	}
	fmt.Printf("%v", getResp.Kvs)
	time.Sleep(20 * time.Second)
}
