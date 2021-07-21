package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	_ "regexp"
	"strings"
)

type KafaConsumerEvent interface {
	KafaMessageNotify(topic, key, message string)
	KafaExceptionNotify(err error)
}

func consumerRun(url string, topic string, notify KafaConsumerEvent) {
	fmt.Println(topic+"-->start kafka consumer monitor......")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2
	config.Net.SASL.Enable = true
	config.Net.SASL.Password = "admin-secret"
	config.Net.SASL.User = "admin"

	client, err := sarama.NewClient(strings.Split(url, ","), config)
	if err != nil {
		if notify != nil {
			go notify.KafaExceptionNotify(err)
		}
		return
	}
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		if notify != nil {
			go notify.KafaExceptionNotify(err)
		}
		return
	}

	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		if notify != nil {
			go notify.KafaExceptionNotify(err)
		}
		return
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			if notify != nil {
				go notify.KafaMessageNotify(msg.Topic, string(msg.Key), string(msg.Value))
			}
		case err := <-partitionConsumer.Errors():
			if notify != nil {
				go notify.KafaExceptionNotify(err)
			}
		}
	}
}

func StartConsumerMonitor(url string, topic string, notify KafaConsumerEvent) {
	go func() {
		consumerRun(url, topic, notify)
	}()
}
