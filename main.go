package main

import (
	"fmt"

	"github.com/zhangjunfang/nsq_manage/nsq_sync"
	"github.com/zhangjunfang/nsq_manage/nsq_sync_config"
)

func main() {
	T()
}
func T() {
	go nsq_sync_config.Consumer()
	//nsq_sync_config.Consumer()
	fmt.Println("-----------------------------")
	nsq_sync_config.Producer()
}

func Sysn() {
	go nsq_sync.StartConsumer()
	fmt.Println("-----------------------------")
	nsq_sync.StartProducer()
}
