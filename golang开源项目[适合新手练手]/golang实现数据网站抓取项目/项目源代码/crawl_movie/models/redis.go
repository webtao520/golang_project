package models

import (
	"github.com/astaxie/goredis"
)

const (
	URL_QUEUE = "url_queue"
	URL_VISIT_SET = "url_visit_set"
)

var (
	client goredis.Client
)

func ConnectRedis(addr string){
	client.Addr = addr
}

func PutinQueue(url string){
	client.Lpush(URL_QUEUE, []byte(url))
}

func PopfromQueue() string{
	res,err := client.Rpop(URL_QUEUE)
	if err != nil{
		panic(err)
	}

	return string(res)
}

func GetQueueLength() int{
	length,err := client.Llen(URL_QUEUE)
	if err != nil{
		return 0
	}

	return length
}

func AddToSet(url string){
	client.Sadd(URL_VISIT_SET, []byte(url))
}

func IsVisit(url string) bool{
	bIsVisit, err := client.Sismember(URL_VISIT_SET, []byte(url))
	if err != nil{
		return false
	}

	return bIsVisit
}

