package main

import (
	grpc_server "book-service/internal/controller/grpc"
	"book-service/internal/database"
	"book-service/internal/repository/mysql"
	"book-service/internal/service/book"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"pkg-service/book_gen"
	"pkg-service/discovery"
	"pkg-service/discovery/consul"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serviceName = "book"
const port = 8081

func main() {
	database, err := database.InitDatabase()

	if err != nil {
		panic(err)
	}

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
	repository := mysql.New(database)
	service := book.New(repository)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("book service started")

	serverPort := fmt.Sprintf("localhost:%d", port)

	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	book_gen.RegisterBookServiceServer(s, grpc_server.New(service))

	reflection.Register(s)

	go func() {
		fmt.Println("Starting book service...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stopping the book service..")
	s.Stop()
	fmt.Println("Stopping listener...")
	lis.Close()
	fmt.Println("End of Program")
}
