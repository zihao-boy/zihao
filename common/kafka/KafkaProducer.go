package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	config2 "github.com/zihao-boy/zihao/config"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func (self *KafkaProducer) Init() error {
	//这里可以初始化多个kafka的,因为是集群，最好多传几个，但是只传一个也可以使用
	servers := []string{fmt.Sprintf("%s:%d", config2.G_AppConfig.KafkaIp, config2.G_AppConfig.KafkaPort)}
	p, err := sarama.NewSyncProducer(servers, sarama.NewConfig())
	if err != nil {
		return err
	}
	self.producer = p
	return nil
}

func (self *KafkaProducer) SendMessage(topic string, data []byte) error {
	if self.producer == nil {
		return errors.New("no producer while send message")
	}
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   nil,
		Value: sarama.ByteEncoder(data),
	}
	_, _, err := self.producer.SendMessage(kafkaMsg)
	return err
}

func (self *KafkaProducer) Close() error {
	if self.producer != nil {
		return self.producer.Close()
	}
	return nil
}
