package kafka_listener

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ProcessMessageCreateFileAnime(msg *kafka.Message) {
	fmt.Println("[KAFKA_TOPIC_DEMO_CREATE] messange: ", string(msg.Value))
	// TODO
}

func ProcessMessageUpdateFileAnime(msg *kafka.Message) {
	fmt.Println("[KAFKA_TOPIC_DEMO_UPDATE] messange: ", string(msg.Value))
	// TODO
}
