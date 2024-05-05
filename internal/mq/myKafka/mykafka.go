package myKafka

import "github.com/segmentio/kafka-go"

func InitKafka() (err error) {
	// 初始化kafka连接
	config := &kafka.ConnConfig{
		ClientID:        1,
		Topic:           "",
		Partition:       0,
		Broker:          0,
		Rack:            "",
		TransactionalID: "",
	}
	return
}
