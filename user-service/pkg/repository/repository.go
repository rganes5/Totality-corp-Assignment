package repository

import (
	"context"
	"errors"
	"user-service/pkg/models"
	"user-service/pkg/repository/interfaces"
)

type userRepo struct {
	UserData map[int32]models.UserData
}

func NewUserRepo() interfaces.UserRepo {
	repository := &userRepo{
		UserData: make(map[int32]models.UserData),
	}
	demoUsers := []models.UserData{
		{
			UserId:    1,
			FirstName: "Ganesh",
			City:      "Kumily",
			Phone:     "9746226152",
			Height:    175,
			Married:   true,
		},
		{
			UserId:    2,
			FirstName: "Stebin",
			City:      "Kottayam",
			Phone:     "9988776655",
			Height:    165,
			Married:   true,
		},
		{
			UserId:    3,
			FirstName: "Sara",
			City:      "Thrissur",
			Phone:     "9876543210",
			Height:    160,
			Married:   true,
		},
		{
			UserId:    4,
			FirstName: "Amit",
			City:      "Delhi",
			Phone:     "1234567890",
			Height:    180,
			Married:   true,
		},
		{
			UserId:    10,
			FirstName: "Kumar",
			City:      "Mumbai",
			Phone:     "9192939495",
			Height:    172,
			Married:   true,
		},
		{
			UserId:    5,
			FirstName: "Emily",
			City:      "New York",
			Phone:     "5551234567",
			Height:    170,
			Married:   true,
		},
		{
			UserId:    6,
			FirstName: "John",
			City:      "London",
			Phone:     "4488996655",
			Height:    175,
			Married:   true,
		},
		{
			UserId:    7,
			FirstName: "Alice",
			City:      "Paris",
			Phone:     "3311223344",
			Height:    162,
			Married:   true,
		},
		{
			UserId:    8,
			FirstName: "Carlos",
			City:      "Madrid",
			Phone:     "3499887766",
			Height:    178,
			Married:   true,
		},
		{
			UserId:    9,
			FirstName: "Maria",
			City:      "Barcelona",
			Phone:     "3466558877",
			Height:    166,
			Married:   true,
		},
	}

	for _, users := range demoUsers {
		repository.UserData[users.UserId] = users
	}
	return repository
}

func (repo *userRepo) GetUserData(ctx context.Context, id int32) (models.UserData, error) {
	users, ok := repo.UserData[id]
	if !ok {
		return models.UserData{}, errors.New("user does not exists")
	}
	return users, nil
}

func (repo *userRepo) GetUserDataList(ctx context.Context, ids []int32) ([]models.UserData, models.NotFoundList, error) {
	var FoundUsers []models.UserData
	var NotFoundUsers models.NotFoundList
	for _, id := range ids {
		users, ok := repo.UserData[id]
		if !ok {
			NotFoundUsers.UsersNotFound = append(NotFoundUsers.UsersNotFound, id)
		} else {
			FoundUsers = append(FoundUsers, users)
		}
	}
	return FoundUsers, NotFoundUsers, nil
}
