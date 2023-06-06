package grpc

import (
	"context"
	"pkg-service/rent_gen"
	"rent-service/internal/service/rent"
	"rent-service/pkg/model"
	"strconv"
)

type Server struct {
	rent_gen.UnimplementedRentServiceServer
	service *rent.Service
}

func New(service *rent.Service) *Server {
	return &Server{
		service: service,
	}
}

func (ctrl *Server) GetAllRents(request *rent_gen.GetAllRentsRequest, stream rent_gen.RentService_GetAllRentsServer) error {
	userId := strconv.Itoa(int(request.GetUserId()))

	rents, err := ctrl.service.GetAll(userId)

	if err != nil {
		return err
	}

	for _, rent := range rents {
		stream.Send(&rent_gen.GetAllRentsResponse{
			Rent: &rent_gen.Rent{
				Id:        uint32(rent.ID),
				UserId:    int32(rent.UserID),
				BookId:    int32(rent.BookID),
				BookTitle: rent.BookTitle,
			},
		})
	}

	return nil
}

func (ctrl *Server) CreateRent(ctx context.Context, request *rent_gen.CreateRentRequest) (*rent_gen.CreateRentResponse, error) {
	var rentInput model.RentInput = model.RentInput{
		UserID:    int(request.GetUserId()),
		BookID:    int(request.GetBookId()),
		BookTitle: request.GetBookTitle(),
	}

	rent, err := ctrl.service.Create(rentInput)

	if err != nil {
		return &rent_gen.CreateRentResponse{
			Status:  "failed",
			Message: "failed to create a book rent",
		}, err
	}

	return &rent_gen.CreateRentResponse{
		Status:  "success",
		Message: "book rent created",
		Rent: &rent_gen.Rent{
			Id:        uint32(rent.ID),
			UserId:    int32(rent.UserID),
			BookId:    int32(rent.BookID),
			BookTitle: rent.BookTitle,
		},
	}, nil
}
