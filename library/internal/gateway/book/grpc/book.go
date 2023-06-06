package grpc

import (
	bookmodel "book-service/pkg/model"
	"context"
	"io"
	grpc_util "library-service/pkg/util"
	"log"
	"pkg-service/discovery"
	"pkg-service/model"
	"pkg-service/proto_gen"
	"pkg-service/util"
)

type Gateway struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
	return &Gateway{
		registry: registry,
	}
}

func (g *Gateway) GetAll(ctx context.Context) (*model.Response[[]bookmodel.Book], error) {
	log.Println("calling book service with gRPC: get all")

	conn, err := util.ConnectgRPCService(ctx, "book", g.registry)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	var client proto_gen.BookServiceClient = proto_gen.NewBookServiceClient(conn)

	stream, err := client.GetAllBooks(ctx, &proto_gen.GetAllBooksRequest{})

	if err != nil {
		return nil, err
	}

	var books []bookmodel.Book

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error when streaming: %v\n", err)
		}

		book := grpc_util.MapToBookModel(res.Book)

		books = append(books, book)
	}

	return &model.Response[[]bookmodel.Book]{
		Status:  "success",
		Message: "all books",
		Data:    books,
	}, nil
}

func (g *Gateway) GetByID(ctx context.Context, id string) (*model.Response[bookmodel.Book], error) {
	log.Println("calling book service with gRPC: get by id")

	conn, err := util.ConnectgRPCService(ctx, "book", g.registry)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	var client proto_gen.BookServiceClient = proto_gen.NewBookServiceClient(conn)

	res, err := client.GetBookByID(ctx, &proto_gen.GetBookByIDRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return &model.Response[bookmodel.Book]{
		Status:  "success",
		Message: "book found",
		Data:    grpc_util.MapToBookModel(res.Book),
	}, nil
}

func (g *Gateway) Create(ctx context.Context, bookInput bookmodel.BookInput) (*model.Response[bookmodel.Book], error) {
	log.Println("calling book service with gRPC: create")

	conn, err := util.ConnectgRPCService(ctx, "book", g.registry)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	var client proto_gen.BookServiceClient = proto_gen.NewBookServiceClient(conn)

	request := &proto_gen.CreateBookRequest{
		Title:       bookInput.Title,
		Description: bookInput.Description,
		Author:      bookInput.Author,
	}

	res, err := client.CreateBook(ctx, request)

	if err != nil {
		return nil, err
	}

	return &model.Response[bookmodel.Book]{
		Status:  "success",
		Message: "book created",
		Data:    grpc_util.MapToBookModel(res.Book),
	}, nil
}
