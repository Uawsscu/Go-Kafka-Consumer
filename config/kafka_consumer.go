package config

import (
	"context"
	"fmt"
	kafka_listener "main/kafka-listener"
	"os"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartKafkaConsumer() {
	var wg sync.WaitGroup
	createKafkaConsumer(os.Getenv("KAFKA_TOPIC_DEMO_CREATE"), &wg, kafka_listener.ProcessMessageCreateFileAnime)
	createKafkaConsumer(os.Getenv("KAFKA_TOPIC_DEMO_UPDATE"), &wg, kafka_listener.ProcessMessageUpdateFileAnime)
	wg.Wait()
}

func createKafkaConsumer(topic string, wg *sync.WaitGroup, messageHandler func(*kafka.Message)) {
	config := kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKERS"),
		"group.id":          "my-consumer-group",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(&config)
	if err != nil {
		fmt.Println("[topic:", topic, "] Error creating Kafka Consumer!! ")
		fmt.Println(err)
		return
	}

	// check topic exist and create topic not exist
	if err := createKafkaTopic(topic); err != nil {
		fmt.Printf("Error creating Kafka topic: %v\n", err)
		return
	}

	// Subscribe topic
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Println("[topic:", topic, "] Error subscribe topic!! ")
		fmt.Println(err)
		return
	}

	// start listener Kafka messages
	go func() {
		defer consumer.Close()
		defer wg.Done()

		fmt.Printf("Listening for messages on topic %s...\n", topic)

		for {
			msg, err := consumer.ReadMessage(-1)
			if err == nil {
				messageHandler(msg)
			} else {
				fmt.Printf("Error reading message for topic %s: %v\n", topic, err)
			}
		}
	}()

	wg.Add(1)
}

func createKafkaTopic(topic string) error {
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_BROKERS")})
	if err != nil {
		return err
	}
	defer adminClient.Close()

	topicSpec := kafka.TopicSpecification{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}

	topics := []kafka.TopicSpecification{topicSpec}
	_, err = adminClient.CreateTopics(context.Background(), topics)
	if err != nil {
		return err
	}

	return nil
}
