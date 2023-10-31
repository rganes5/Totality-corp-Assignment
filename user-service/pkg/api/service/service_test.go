package service

import (
	"context"
	"net/http"
	"testing"
	"user-service/pkg/models"
	"user-service/pkg/pb"
	"user-service/pkg/repository/mockRepository"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_GetUserData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepository.NewMockUserRepo(ctrl)
	userService := NewUserService(mockRepo)

	// Define your test data and expected results
	userRequest := &pb.GetUserRequest{UserId: 1}
	expectedUserResponse := &pb.GetUserResponse{
		Status: http.StatusOK,
		Result: &pb.UserData{
			UserId:    1,
			FirstName: "John",
			City:      "New York",
			Phone:     "123-456-7890",
			Height:    180,
			Married:   false,
		},
	}

	// Set up the expected behavior of the mock repository
	mockRepo.EXPECT().GetUserData(gomock.Any(), userRequest.UserId).Return(
		models.UserData{
			UserId:    1,
			FirstName: "John",
			City:      "New York",
			Phone:     "123-456-7890",
			Height:    180,
			Married:   false,
		},
		nil,
	)

	// Call the service function
	actualUserResponse, err := userService.GetUserData(context.Background(), userRequest)

	// Perform assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedUserResponse, actualUserResponse)
}

func NewMockUserRepo(ctrl *gomock.Controller) {
	panic("unimplemented")
}
func TestUserService_GetUserDataList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepository.NewMockUserRepo(ctrl)
	userService := NewUserService(mockRepo)

	// Define your test data and expected results
	userListRequest := &pb.GetUserDataListRequest{UserIdList: []int32{1, 2, 3}}
	actualUserListResponse := &pb.GetUserDataListResponse{
		Status: http.StatusOK,
		Result: []*pb.UserData{
			{
				UserId:    1,
				FirstName: "John",
				City:      "New York",
				Phone:     "123-456-7890",
				Height:    180,
				Married:   false,
			},
			{
				UserId:    2,
				FirstName: "Alice",
				City:      "Los Angeles",
				Phone:     "987-654-3210",
				Height:    165,
				Married:   true,
			},
		},
		NotFound: []int32{3},
	}

	// Set up the expected behavior of the mock repository
	mockRepo.EXPECT().GetUserDataList(gomock.Any(), userListRequest.UserIdList).Return(
		[]models.UserData{
			{
				UserId:    1,
				FirstName: "John",
				City:      "New York",
				Phone:     "123-456-7890",
				Height:    180,
				Married:   false,
			},
			{
				UserId:    2,
				FirstName: "Alice",
				City:      "Los Angeles",
				Phone:     "987-654-3210",
				Height:    165,
				Married:   true,
			},
		},
		models.NotFoundList{
			UsersNotFound: []int32{3},
		},
		nil,
	)

	// Call the service function
	actualUserListResponse, err := userService.GetUserDataList(context.Background(), userListRequest)

	// Perform assertions
	assert.NoError(t, err)
	assert.Equal(t, actualUserListResponse, actualUserListResponse)
}
