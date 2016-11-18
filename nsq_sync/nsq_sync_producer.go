package nsq_sync

import (
	"log"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
)

// 生产者
func StartProducer() {
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
