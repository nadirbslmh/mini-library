package util

import (
	"pkg-service/proto_gen"
	"rent-service/pkg/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapToRentPb(rent model.Rent) *proto_gen.Rent {
	return &proto_gen.Rent{
		Id:        uint32(rent.ID),
		CreatedAt: timestamppb.New(rent.CreatedAt),
		UpdatedAt: timestamppb.New(rent.UpdatedAt),
		DeletedAt: &proto_gen.DeletedAt{
			Time:  timestamppb.New(rent.DeletedAt.Time),
			Valid: rent.DeletedAt.Valid,
		},
		UserId:    int32(rent.UserID),
		BookId:    int32(rent.BookID),
		BookTitle: rent.BookTitle,
	}
}
