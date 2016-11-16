package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	go startConsumer()
	fmt.Println("-----------------------------")
	startProducer()
}

// 生产者
func startProducer() {
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer("192.168.1.90:4150", cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 发布消息
	var i int64 = 0
	for {
		i++
		if err := producer.Publish("test", []byte("test message"+strconv.Itoa(int(i)))); err != nil {
			log.Fatal("publish error: " + err.Error())
		}
		time.Sleep(0 * time.Second)
	}
}

// 消费者
func startConsumer() {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "sensor01", cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 设置消息处理函数
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(string(message.Body))
		return nil
	}))
	// 连接到单例nsqd
	if err := consumer.ConnectToNSQD("192.168.1.90:4150"); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
}
