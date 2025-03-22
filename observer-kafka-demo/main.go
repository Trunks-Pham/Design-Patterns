package main

import (
	"fmt"
	"observer-kafka-demo/consumer"
	"observer-kafka-demo/producer"
	"time"
)

func main() {
	broker := "localhost:9092"
	topic := "observer-demo"

	// Tạo các observer (consumer)
	consumer.StartConsumer(broker, topic, "group-1", "Observer A")
	consumer.StartConsumer(broker, topic, "group-2", "Observer B")

	// Delay khởi động để đảm bảo consumer đã kết nối
	time.Sleep(2 * time.Second)

	// Gửi message (act như Subject)
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Event #%d", i+1)
		producer.PublishMessage(broker, topic, msg)
		time.Sleep(1 * time.Second)
	}

	// Đợi để consumer xử lý
	time.Sleep(5 * time.Second)
}
