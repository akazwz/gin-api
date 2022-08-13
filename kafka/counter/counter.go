package counter

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func Produce() {
	mechanism, err := scram.Mechanism(scram.SHA256, os.Getenv("KAFKA_USERNAME"), os.Getenv("KAFKA_PASSWORD"))

	sharedTransport := &kafka.Transport{
		SASL: mechanism,
		TLS:  &tls.Config{},
	}

	_ = kafka.Client{
		Addr:      kafka.TCP(os.Getenv("KAFKA_ADDR")),
		Timeout:   10 * time.Second,
		Transport: sharedTransport,
	}

	w := kafka.Writer{
		Addr:      kafka.TCP("honest-trout-9670-us1-kafka.upstash.io:9092"),
		Topic:     "counter",
		Balancer:  &kafka.LeastBytes{},
		Transport: sharedTransport,
	}

	err = w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("post_view"),
			Value: []byte("10"),
		},
	)
	if err != nil {
		log.Fatalln("failed to write messages:", err)
	}
}

func Consume() {
	mechanism, err := scram.Mechanism(scram.SHA256, os.Getenv("KAFKA_USERNAME"), os.Getenv("KAFKA_PASSWORD"))
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	r := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers: []string{"honest-trout-9670-us1-kafka.upstash.io:9092"},
			Topic:   "counter",
			Dialer:  dialer,
		},
	)
	message, err := r.ReadMessage(context.Background())
	if err != nil {
		log.Fatalln("fetch message error:", err)
		return
	}
	log.Println("key:", string(message.Key))
	log.Println("value:", string(message.Value))
}
