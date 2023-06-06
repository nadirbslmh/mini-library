package grpc

import (
	"context"
	"errors"
	"io"
	bookgateway "library-service/internal/gateway/book/grpc"
	"library-service/pkg/constant"
	"log"
	"pkg-service/discovery"
	"pkg-service/model"
	"pkg-service/rent_gen"
	"pkg-service/util"
	rentmodel "rent-service/pkg/model"
	"strconv"
)

type Gateway struct {
	bookgateway *bookgateway.Gateway
	registry    discovery.Registry
}

func New(registry discovery.Registry, bookgateway *bookgateway.Gateway) *Gateway {
	return &Gateway{
		registry:    registry,
		bookgateway: bookgateway,
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

	var client rent_gen.RentServiceClient = rent_gen.NewRentServiceClient(conn)

	request := &rent_gen.GetAllRentsRequest{
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

		rent := rentmodel.Rent{
			ID:        uint(res.Rent.Id),
			UserID:    int(res.Rent.UserId),
			BookID:    int(res.Rent.BookId),
			BookTitle: res.Rent.BookTitle,
		}

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

	var client rent_gen.RentServiceClient = rent_gen.NewRentServiceClient(conn)

	request := &rent_gen.CreateRentRequest{
		UserId:    int32(rentInput.UserID),
		BookId:    int32(rentInput.BookID),
		BookTitle: rentInput.BookTitle,
	}

	res, err := client.CreateRent(ctx, request)

	if err != nil {
		return nil, err
	}

	return &model.Response[rentmodel.Rent]{
		Status:  "success",
		Message: "book rent created",
		Data: rentmodel.Rent{
			ID:        uint(res.Rent.Id),
			UserID:    int(res.Rent.UserId),
			BookID:    int(res.Rent.BookId),
			BookTitle: res.Rent.BookTitle,
		},
	}, nil
}
