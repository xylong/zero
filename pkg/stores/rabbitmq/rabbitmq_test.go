package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"testing"
)

func TestPublish(t *testing.T) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", "root", "123456", "127.0.0.1", 5672)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	//队列创建成功
	queue, err := c.QueueDeclare("test", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = c.Publish("", queue.Name, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("test001"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("发送消息成功")
}

func TestConsumer(t *testing.T) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", "root", "123456", "127.0.0.1", 5672)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	msgs, err := c.Consume("test", "c1", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	for msg := range msgs {
		msg.Ack(true)
		fmt.Println(msg.DeliveryTag, string(msg.Body))
	}
}
