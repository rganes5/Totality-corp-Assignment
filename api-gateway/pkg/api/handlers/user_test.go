package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"api-gateway/pkg/pb"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockUserClient struct {
	mockGetUserData     func(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error)
	mockGetUserDataList func(ctx context.Context, req *pb.GetUserDataListRequest) (*pb.GetUserDataListResponse, error)
}

func (m *mockUserClient) GetUserData(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if m.mockGetUserData != nil {
		return m.mockGetUserData(ctx, req)
	}
	return nil, errors.New("Not implemented")
}

func (m *mockUserClient) GetUserDataList(ctx context.Context, req *pb.GetUserDataListRequest) (*pb.GetUserDataListResponse, error) {
	if m.mockGetUserDataList != nil {
		return m.mockGetUserDataList(ctx, req)
	}
	return nil, errors.New("Not implemented")
}

func TestUserHandler_GetUserData(t *testing.T) {
	handler := NewUserHandler(&mockUserClient{
		mockGetUserData: func(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
			if req.UserId == 1 {
				return &pb.GetUserResponse{
					Status: http.StatusOK,
					Result: &pb.UserData{
						UserId:    1,
						FirstName: "John",
						City:      "New York",
						Phone:     "123-456-7890",
						Height:    180,
						Married:   false,
					},
				}, nil
			}
			return nil, errors.New("User not found")
		},
	})

	r := gin.Default()
	r.GET("/user/getbyid/:userId", handler.GetUserData)
	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/user/getbyid/1")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUserHandler_GetUserDataList(t *testing.T) {
	handler := NewUserHandler(&mockUserClient{
		mockGetUserDataList: func(ctx context.Context, req *pb.GetUserDataListRequest) (*pb.GetUserDataListResponse, error) {
			if len(req.UserIdList) > 0 {
				return &pb.GetUserDataListResponse{
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
					},
					NotFound: []int32{3},
				}, nil
			}
			return nil, errors.New("No users found")
		},
	})

	r := gin.Default()
	r.POST("/user/getbylist", handler.GetUserDataList)
	ts := httptest.NewServer(r)
	defer ts.Close()

	payload := `{"userId": [1, 2, 3]}`
	resp, err := http.Post(ts.URL+"/user/getbylist", "application/json", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
