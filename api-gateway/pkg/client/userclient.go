package client

import (
	"context"

	"api-gateway/pkg/client/interfaces"
	"api-gateway/pkg/pb"
	"api-gateway/pkg/service"
)

type userClient struct {
	Client pb.UserServiceClient
}

func NewUserClient(service service.Clients) interfaces.UserClient {
	return &userClient{
		Client: service.Usercli,
	}
}

func (cr *userClient) GetUserData(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	response, err := cr.Client.GetUserData(ctx, request)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (cr *userClient) GetUserDataList(ctx context.Context, request *pb.GetUserDataListRequest) (*pb.GetUserDataListResponse, error) {
	response, err := cr.Client.GetUserDataList(ctx, request)
	if err != nil {
		return response, err
	}
	return response, nil
}
