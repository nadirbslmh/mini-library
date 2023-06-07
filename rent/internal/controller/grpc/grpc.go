package grpc

import (
	"context"
	"pkg-service/proto_gen"
	"rent-service/internal/service/rent"
	"rent-service/pkg/model"
	"rent-service/pkg/util"
)

type Server struct {
	proto_gen.UnimplementedRentServiceServer
	service *rent.Service
}

func New(service *rent.Service) *Server {
	return &Server{
		service: service,
	}
}

func (ctrl *Server) GetAllRents(request *proto_gen.GetAllRentsRequest, stream proto_gen.RentService_GetAllRentsServer) error {
	userId := request.GetUserId()

	rents, err := ctrl.service.GetAll(userId)

	if err != nil {
		return err
	}

	for _, rent := range rents {
		stream.Send(&proto_gen.GetAllRentsResponse{
			Rent: util.MapToRentPb(rent),
		})
	}

	return nil
}

func (ctrl *Server) CreateRent(ctx context.Context, request *proto_gen.CreateRentRequest) (*proto_gen.CreateRentResponse, error) {
	var rentInput model.RentInput = model.RentInput{
		UserID:    int(request.GetUserId()),
		BookID:    int(request.GetBookId()),
		BookTitle: request.GetBookTitle(),
	}

	err := rentInput.Validate()

	if err != nil {
		return &proto_gen.CreateRentResponse{
			Status:  "failed",
			Message: "validation failed",
		}, err
	}

	rent, err := ctrl.service.Create(rentInput)

	if err != nil {
		return &proto_gen.CreateRentResponse{
			Status:  "failed",
			Message: "failed to create a book rent",
		}, err
	}

	return &proto_gen.CreateRentResponse{
		Status:  "success",
		Message: "book rent created",
		Rent:    util.MapToRentPb(rent),
	}, nil
}
