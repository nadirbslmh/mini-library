package main

import (
	"context"
	"fmt"
	"log"
	"logging-service/internal/database"
	"logging-service/internal/routes"

	"pkg-service/discovery"
	"pkg-service/discovery/consul"

	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const serviceName = "logging"
const port = 8085

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	database, err := database.InitMongo("library")

	if err != nil {
		panic(err)
	}

	log.Println("connected to the mongoDB")

	routes.SetupRoutes(e, database)

	// start registry
	registry, err := consul.NewRegistry("consul-service:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("logging-service:%d", port)); err != nil {
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

	// start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
