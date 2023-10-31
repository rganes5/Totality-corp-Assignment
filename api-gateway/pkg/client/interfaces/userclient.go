package interfaces

import (
	"context"

	"api-gateway/pkg/pb"
)

type UserClient interface {
	GetUserData(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error)
	GetUserDataList(ctx context.Context, request *pb.GetUserDataListRequest) (*pb.GetUserDataListResponse, error)
}
