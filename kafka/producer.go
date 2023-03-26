package kafka

import (
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func NewKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func CreateTopic(kafkaURL, topic string) {

	conn, err := kafka.Dial("tcp", kafkaURL)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     12,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}

// func main() {
// 	kafkaURL := "localhost:9092"
// 	topic := "quickstart"

// 	writer := newKafkaWriter(kafkaURL, topic)
// 	defer writer.Close()
// 	fmt.Println("start producing ... !!")
// 	for i := 0; ; i++ {
// 		key := fmt.Sprintf("Key-%d", i)
// 		msg := kafka.Message{
// 			Value: []byte(fmt.Sprintf("name-%d", i)),
// 		}
// 		err := writer.WriteMessages(context.Background(), msg)
// 		if err != nil {
// 			fmt.Println(err)
// 		} else {
// 			fmt.Println("produced", key)
// 		}
// 		time.Sleep(1 * time.Second)
// 	}
// }
