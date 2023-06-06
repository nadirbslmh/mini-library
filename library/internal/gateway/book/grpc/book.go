package grpc

import (
	bookmodel "book-service/pkg/model"
	"context"
	"io"
	"log"
	"pkg-service/book_gen"
	"pkg-service/discovery"
	"pkg-service/model"
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

	var client book_gen.BookServiceClient = book_gen.NewBookServiceClient(conn)

	stream, err := client.GetAllBooks(ctx, &book_gen.GetAllBooksRequest{})

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

		book := bookmodel.Book{
			ID:          uint(res.Book.Id),
			Title:       res.Book.Title,
			Description: res.Book.Description,
			Author:      res.Book.Author,
		}

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

	var client book_gen.BookServiceClient = book_gen.NewBookServiceClient(conn)

	res, err := client.GetBookByID(ctx, &book_gen.GetBookByIDRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return &model.Response[bookmodel.Book]{
		Status:  "success",
		Message: "book found",
		Data: bookmodel.Book{
			ID:          uint(res.Book.Id),
			Title:       res.Book.Title,
			Description: res.Book.Description,
			Author:      res.Book.Author,
		},
	}, nil
}

func (g *Gateway) Create(ctx context.Context, bookInput bookmodel.BookInput) (*model.Response[bookmodel.Book], error) {
	log.Println("calling book service with gRPC: create")

	conn, err := util.ConnectgRPCService(ctx, "book", g.registry)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	var client book_gen.BookServiceClient = book_gen.NewBookServiceClient(conn)

	request := &book_gen.CreateBookRequest{
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
		Data: bookmodel.Book{
			ID:          uint(res.Book.Id),
			Title:       res.Book.Title,
			Description: res.Book.Description,
			Author:      res.Book.Author,
		},
	}, nil
}
