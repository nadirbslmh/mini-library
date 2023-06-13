package main

import (
	"context"
	"fmt"
	"listener/pkg/util"
	"log"
	"os"
	"os/signal"
	"pkg-service/constant"
	"pkg-service/discovery"
	"pkg-service/discovery/consul"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const serviceName = "listener"
const port = 9000

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "kafka-service:9092",
		"group.id":          "kafka-go",
		"auto.offset.reset": "earliest",
	}

	consumer, err := util.CreateConsumer(config)

	if err != nil {
		panic(err)
	}

	logTopic := constant.LOG_TOPIC

	topics := []string{logTopic}

	err = consumer.SubscribeTopics(topics, nil)

	if err != nil {
		panic(err)
	}

	// start registry
	registry, err := consul.NewRegistry("consul-service:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("listener-service:%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	log.Println("listener is started")

	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			util.HandleMessage(string(ev.Key), ev.Value, registry)
		}
	}

	consumer.Close()
}
