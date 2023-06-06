package main

import (
	grpc_server "auth-service/internal/controller/grpc"
	"auth-service/internal/database"
	"auth-service/internal/repository/mysql"
	"auth-service/internal/service/auth"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"pkg-service/discovery"
	"pkg-service/discovery/consul"
	"pkg-service/proto_gen"

	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serviceName = "auth"
const port = 8083

func main() {
	database, err := database.InitDatabase()

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
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("auth-service:%d", port)); err != nil {
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
	repository := mysql.New(database)
	service := auth.New(repository)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("auth service started")

	serverPort := fmt.Sprintf("auth-service:%d", port)

	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	proto_gen.RegisterAuthServiceServer(s, grpc_server.New(service))

	reflection.Register(s)

	go func() {
		fmt.Println("Starting auth service...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stopping the auth service..")
	s.Stop()
	fmt.Println("Stopping listener...")
	lis.Close()
	fmt.Println("End of Program")
}
