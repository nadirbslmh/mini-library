package main

import (
	"auth-service/internal/database"
	"auth-service/internal/routes"
	"context"
	"fmt"
	"log"

	"pkg-service/discovery"
	"pkg-service/discovery/consul"

	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const serviceName = "auth"
const port = 8083

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	database, err := database.InitDatabase()

	if err != nil {
		panic(err)
	}

	routes.SetupRoutes(e, database)

	// start registry
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
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
