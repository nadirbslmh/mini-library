package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	bookgateway "library-service/internal/gateway/book/grpc"
	"library-service/pkg/constant"
	app_util "library-service/pkg/util"
	"log"
	app_constant "pkg-service/constant"
	"pkg-service/discovery"
	"pkg-service/model"
	"pkg-service/proto_gen"
	"pkg-service/util"
	rentmodel "rent-service/pkg/model"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Gateway struct {
	bookgateway *bookgateway.Gateway
	registry    discovery.Registry
	producer    *kafka.Producer
}

func New(registry discovery.Registry, bookgateway *bookgateway.Gateway, producer *kafka.Producer) *Gateway {
	return &Gateway{
		registry:    registry,
		bookgateway: bookgateway,
		producer:    producer,
	}
}

func (g *Gateway) GetAll(ctx context.Context) (*model.Response[[]rentmodel.Rent], error) {
	log.Println("calling rent service with gRPC: get all")

	userCtx := ctx.Value(constant.USER_ID_KEY)

	userID, ok := userCtx.(string)

	if !ok {
		return nil, errors.New("id is invalid")
	}

	conn, err := util.ConnectgRPCService(ctx, "rent", g.registry)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	var client proto_gen.RentServiceClient = proto_gen.NewRentServiceClient(conn)

	request := &proto_gen.GetAllRentsRequest{
		UserId: userID,
	}

	stream, err := client.GetAllRents(ctx, request)

	if err != nil {
		return nil, err
	}

	var rents []rentmodel.Rent

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error when streaming: %v\n", err)
		}

		rent := app_util.MapToRentModel(res.Rent)

		rents = append(rents, rent)
	}

	return &model.Response[[]rentmodel.Rent]{
		Status:  "success",
		Message: "all book rents",
		Data:    rents,
	}, nil
}

func (g *Gateway) Create(ctx context.Context, rentInput rentmodel.RentInput) (*model.Response[rentmodel.Rent], error) {
	log.Println("calling rent service with gRPC: create")

	responseData, err := g.bookgateway.GetByID(ctx, strconv.Itoa(rentInput.BookID))

	isFailed := err != nil || responseData == nil

	if isFailed {
		return nil, err
	}

	rentedBook := responseData.Data

	rentInput.BookTitle = rentedBook.Title

	conn, err := util.ConnectgRPCService(ctx, "rent", g.registry)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	var client proto_gen.RentServiceClient = proto_gen.NewRentServiceClient(conn)

	request := &proto_gen.CreateRentRequest{
		UserId:    int32(rentInput.UserID),
		BookId:    int32(rentInput.BookID),
		BookTitle: rentInput.BookTitle,
	}

	res, err := client.CreateRent(ctx, request)

	if err != nil {
		return nil, err
	}

	//TODO: write rent log with Kafka
	go func() {
		for e := range g.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	rentID := strconv.Itoa(int(res.GetRent().GetId()))

	key := []byte("log" + ":" + rentID)

	input, err := json.Marshal(&rentInput)

	if err != nil {
		return nil, err
	}

	topic := app_constant.LOG_TOPIC

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          input,
	}

	err = app_util.SendMessage(g.producer, message)

	if err != nil {
		return nil, err
	}

	return &model.Response[rentmodel.Rent]{
		Status:  "success",
		Message: "book rent created",
		Data:    app_util.MapToRentModel(res.Rent),
	}, nil
}
