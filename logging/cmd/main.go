package main

import (
	"context"
	"fmt"
	"log"
	"logging-service/internal/database"
	"logging-service/internal/routes"

	"pkg-service/constant"
	"pkg-service/discovery"
	"pkg-service/discovery/consul"
	"pkg-service/util"

	"time"

	"github.com/hashicorp/consul/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const serviceName = "logging"
const port = 8085

func main() {
	// init config
	consulCfg := api.DefaultConfig()
	consulCfg.Address = "localhost:8500"

	client, err := api.NewClient(consulCfg)
	if err != nil {
		panic(err)
	}

	dbName, err := util.GetConfigValue(client, constant.LOG_DB_NAME)
	if err != nil {
		panic(err)
	}

	mongoURI, err := util.GetConfigValue(client, constant.LOG_MONGO_URI)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	database, err := database.InitMongo(dbName, mongoURI)

	if err != nil {
		panic(err)
	}

	log.Println("connected to the mongoDB")

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
