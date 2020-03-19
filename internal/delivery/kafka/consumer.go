package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	//spServiceEntity "github.com/vahnblue/1BService/internal/entity/spService"
	appleEntity "print-apple/internal/entity/apple"
	"print-apple/pkg/kafka"

	"github.com/Shopify/sarama"

	//spServiceEntity "github.com/vahnblue/1BService/internal/entity/spService"
	"log"
)

//IAppleSvc is an interface to User Service
type IAppleSvc interface {
	GetPrintApple(ctx context.Context) ([]appleEntity.Apple, error)
	// InsertMany(ctx context.Context, userList []appleEntity.Apple) error
}

type (
	// Consumer represents a Sarama consumer group consumer
	Consumer struct {
		Ready    chan bool
		appleSvc IAppleSvc
	}
)

// New for bridging product handler initialization
func New(apple1 IAppleSvc, k *kafka.Kafka, subscriptions []string) Consumer {
	consumer := Consumer{
		Ready:    make(chan bool),
		appleSvc: apple1,
	}
	// Consumer
	go func() {
		client := k.GetConsumerGroup()
		ctx := context.Background()

		for {
			if err := client.Consume(ctx, subscriptions, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.Ready = make(chan bool)
		}
	}()

	<-consumer.Ready
	return consumer
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.Ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		log.Printf("[KAFKA][CONSUMER] Message claimed: Value = %s, Timestamp = %v, Topic = %s", string(message.Value), message.Timestamp, message.Topic)
		switch message.Topic {
		case "Print_Queue":
			var data appleEntity.Apple
			err := json.Unmarshal(message.Value, &data)
			if err != nil {
				log.Fatalf(err.Error())
			}
			_, err = consumer.appleSvc.GetPrintApple(context.Background())
			//, data.TransFH
			fmt.Println(err)
			if err != nil {
				log.Fatalf(err.Error())
			}

			session.MarkMessage(message, "")
		}
		session.MarkMessage(message, "")

	}

	return nil
}
