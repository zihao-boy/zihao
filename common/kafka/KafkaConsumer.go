package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	config2 "github.com/zihao-boy/zihao/config"
	"strings"
)

type KafkaConsumer struct {
	consumer *cluster.Consumer
}

func (self *KafkaConsumer) Init() error {
	brokersServers := []string{fmt.Sprintf("%s:%d", config2.G_AppConfig.KafkaIp, config2.G_AppConfig.KafkaPort)}
	config := cluster.NewConfig()
	//配置是否接受错误信息
	config.Consumer.Return.Errors = true
	//配置是否接受注意消息
	config.Group.Return.Notifications = true
	//配置是否接受最新消息
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	//这个消费者是谁，同一个消费者如果对一条信息确认了，则不会重复发送
	config.ClientID = config2.G_AppConfig.KafkaGroup
	//topic是指要收到的消息对象
	cg, err := cluster.NewConsumer(brokersServers, config2.G_AppConfig.KafkaGroup, strings.Split(config2.G_AppConfig.KafkaTopic, ","), config)
	if err != nil {
		return err
	}
	self.consumer = cg
	return nil
}

//注意该方法是非阻塞的，如果调用了该方法，并且没有其他的阻塞方法，记得手动阻塞他
func (self *KafkaConsumer) StartKafkaListen(listenMsg func(*sarama.ConsumerMessage)) error {
	if self.consumer == nil {
		return errors.New("还没初始化消费者对象")
	}
	go func(cg *cluster.Consumer) {
		for message := range cg.Messages() {
			go listenMsg(message)
			//确认这条消息收到
			cg.MarkOffset(message, "")
		}
	}(self.consumer)
	go func(cg *cluster.Consumer) {
		for ntf := range cg.Notifications() {
			glog.Infof("%+v", *ntf)
		}
	}(self.consumer)
	go func(cg *cluster.Consumer) {
		for err := range cg.Errors() {
			glog.Errorf("%+v", err)
		}
	}(self.consumer)
	return nil
}
