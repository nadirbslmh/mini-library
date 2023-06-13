package main

import (
	"context"
	"fmt"
	"library-service/internal/routes"
	"library-service/pkg/util"
	"log"
	"pkg-service/discovery"
	"pkg-service/discovery/consul"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const serviceName = "library"
const port = 8080

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// start registry
	registry, err := consul.NewRegistry("consul-service:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("library-service:%d", port)); err != nil {
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

	// create kafka producer
	config := &kafka.ConfigMap{
		"bootstrap.servers": "kafka-service:9092",
	}

	producer, err := util.CreateProducer(config)

	if err != nil {
		panic(err)
	}

	routes.SetupRoutes(e, registry, producer)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
