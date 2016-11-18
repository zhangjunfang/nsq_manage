package nsq_sync_config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
)

func Producer() {
	p, err := nsq.NewProducer("192.168.1.90:4150", nsq.NewConfig())
	if err != nil {
		return
	}
	for i := 0; i < 1000; i++ {
		p.DeferredPublish("topic", time.Millisecond*1, []byte("dfsfdsfdsf:"+strconv.Itoa(i)))
	}

	//p.Stop()
	fmt.Println("--------------", "Producer finshed")
}
