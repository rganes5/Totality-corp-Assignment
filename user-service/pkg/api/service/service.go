package service

import (
	"context"
	"net/http"
	"user-service/pkg/pb"
	"user-service/pkg/repository/interfaces"
)

type userService struct {
	pb.UnimplementedUserServiceServer
	Repo interfaces.UserRepo
}

func NewUserService(repo interfaces.UserRepo) pb.UserServiceServer {
	return &userService{
		Repo: repo,
	}
}

func (cr *userService) GetUserData(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	response, err := cr.Repo.GetUserData(context.Background(), req.UserId)
	if err != nil {
		return &pb.GetUserResponse{
			Status: http.StatusInternalServerError,
		}, err
	}
	users := &pb.UserData{
		UserId:    response.UserId,
		FirstName: response.FirstName,
		City:      response.City,
		Phone:     response.Phone,
		Height:    response.Height,
		Married:   response.Married,
	}
	return &pb.GetUserResponse{
		Status: http.StatusOK,
		Result: users,
	}, nil
}

func (cr *userService) GetUserDataList(ctx context.Context, req *pb.GetUserDataListRequest) (*pb.GetUserDataListResponse, error) {
	response, notFound, err := cr.Repo.GetUserDataList(context.Background(), req.UserIdList)
	if err != nil {
		return &pb.GetUserDataListResponse{
			Status: http.StatusInternalServerError,
		}, nil
	}

	userDataList := make([]*pb.UserData, 0, len(response))
	for _, users := range response {
		userDataList = append(userDataList, &pb.UserData{
			UserId:    users.UserId,
			FirstName: users.FirstName,
			City:      users.City,
			Phone:     users.Phone,
			Height:    users.Height,
			Married:   users.Married,
		})
	}
	return &pb.GetUserDataListResponse{
		Status:   http.StatusOK,
		Result:   userDataList,
		NotFound: notFound.UsersNotFound,
	}, nil
}
