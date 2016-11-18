package nsq_sync

import (
	"log"

	"github.com/nsqio/go-nsq"
)

// 消费者
func StartConsumer() {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "sensor01", cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 设置消息处理函数
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println("consumer============", string(message.Body))
		return nil
	}))
	// 连接到单例nsqd
	if err := consumer.ConnectToNSQD("192.168.1.90:4150"); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
}
