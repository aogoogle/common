package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
)

type KafkaEvent struct {
	Topic   string `json:"topic"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

type KafkaProducerNotify interface {
	KafkaSendSuccessedNotify(event KafkaEvent)
}

type KafkaProducer struct {
	producer sarama.AsyncProducer
	notify KafkaProducerNotify
}

func InitWithUrl(url string, notify KafkaProducerNotify) (producer *KafkaProducer) {
	fmt.Printf("producer_test\n")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Net.SASL.Enable = true
	config.Net.SASL.Password = "admin-secret"
	config.Net.SASL.User = "admin"
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	kafkaProducer := new(KafkaProducer)
	kafkaProducer.notify = notify
	var err error
	client, err := sarama.NewClient(strings.Split(url, ","), config)
	if err != nil {
		fmt.Printf("kafka producer init exception:", err.Error())
		return
	}

	kafkaProducer.producer, err = sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		fmt.Printf("producer create producer error :%s\n", err.Error())
	}
	return kafkaProducer
}

func (k *KafkaProducer) SendMessage(topic string, key string, message string) {
	if k.producer != nil {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(key),
			Value: sarama.StringEncoder(message),
		}
		fmt.Println(msg)
		k.producer.Input() <- msg

		select {
		case suc := <-k.producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %s", suc.Offset, suc.Timestamp.String())
			if k.notify != nil {
				k.notify.KafkaSendSuccessedNotify(KafkaEvent{Topic: topic, Key: key, Message: message})
			}
		case fail := <-k.producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		default:
			fmt.Printf("default\n")
		}
	}
}

func (k *KafkaProducer) SendEvent(event KafkaEvent) {
	k.SendMessage(event.Topic, event.Key, event.Message)
}
