package nsq_sync_config

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/nsqio/go-nsq"
)

func Consumer() {
	mm := nsq.HandlerFunc(Handler)
	c, err := nsq.NewConsumer("topic", "channel", nsq.NewConfig())
	if err != nil {
		return
	}
	c.AddConcurrentHandlers(mm, 16)
	//c.ConnectToNSQDs([]string{"192.168.1.90:4150", "192.168.1.90:4151"})
	c.ConnectToNSQDs([]string{"192.168.1.90:4150"})
	<-c.StopChan
	//c.Stop()
	fmt.Println("--------------", "Consumer stop")
}
func Handler(m *nsq.Message) error {
	mm := make([]byte, 16)
	for _, v := range m.ID {
		mm = append(mm, v)
	}
	b := []byte(mm)
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	fmt.Println("dsfsdfsdf", x, "---", string(m.Body), "====", m.Attempts)
	return nil
}
