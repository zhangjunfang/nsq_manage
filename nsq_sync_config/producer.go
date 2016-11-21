package nsq_sync_config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
)

var P *nsq.Producer

func Producer() {
	//DeferredPublish()
	//DeferredPublishAsync()
	//MultiPublishAsync()
	//PublishAsync()
	MultiPublish()
}
func MultiPublish() {
	p, err := nsq.NewProducer("192.168.1.90:4150", nsq.NewConfig())
	if err != nil {
		return
	}

	for i := 0; i < 2000; i++ {
		p.MultiPublish("topic", [][]byte{[]byte("sfsdfsdf:" + strconv.Itoa(i)), []byte("i dont known")})
	}
	time.Sleep(time.Second * 4)
	fmt.Println("--------------", "Producer finshed")
}

//丢数据  有待研究
func PublishAsync() {
	p, err := nsq.NewProducer("192.168.1.90:4150", nsq.NewConfig())
	if err != nil {
		return
	}
	doneChan := make(chan *nsq.ProducerTransaction)
	for i := 0; i < 2000; i++ {
		p.PublishAsync("topic", []byte("sfsdfsdf:"+strconv.Itoa(i)), doneChan)
	}
	fmt.Println(<-doneChan)
	time.Sleep(time.Second * 4)
	fmt.Println("--------------", "Producer finshed")
}

//有问题 待研究  同样的现象
func MultiPublishAsync() {
	p, err := nsq.NewProducer("192.168.1.90:4150", nsq.NewConfig())
	if err != nil {
		return
	}
	doneChan := make(chan *nsq.ProducerTransaction)
	for i := 0; i < 2000; i++ {
		p.MultiPublishAsync("topic", [][]byte{[]byte("sfsdfsdf:" + strconv.Itoa(i)), []byte("i dont known")}, doneChan)
	}
	fmt.Println(<-doneChan)
	time.Sleep(time.Second * 4)
	fmt.Println("--------------", "Producer finshed")
}

//有问题 待研究
func DeferredPublishAsync() {
	p, err := nsq.NewProducer("192.168.1.90:4150", nsq.NewConfig())
	if err != nil {
		return
	}
	doneChan := make(chan *nsq.ProducerTransaction)

	for i := 0; i < 2000; i++ {
		p.DeferredPublishAsync("topic", (time.Millisecond * 1), []byte("sfsdfsdf:"+strconv.Itoa(i)), doneChan)
	}
	fmt.Println(<-doneChan)
	time.Sleep(time.Second * 4)
	fmt.Println("--------------", "Producer finshed")
}
func DeferredPublish() {
	p, err := nsq.NewProducer("192.168.1.90:4150", nsq.NewConfig())
	if err != nil {
		return
	}
	for i := 0; i < 2000; i++ {
		p.DeferredPublish("topic", time.Millisecond*1, []byte("dfsfdsfdsf:"+strconv.Itoa(i)))
	}
	time.Sleep(time.Second * 4)
	fmt.Println("--------------", "Producer finshed")
}
func Publish() {
	p, err := nsq.NewProducer("192.168.1.90:4150", nsq.NewConfig())
	if err != nil {
		return
	}
	for i := 0; i < 2000; i++ {
		p.Publish("topic", []byte("dfsfdsfdsf:"+strconv.Itoa(i)))
	}
	time.Sleep(time.Second * 4)
	fmt.Println("--------------", "Producer finshed")
}
