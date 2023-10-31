package interfaces

import (
	"context"
	"user-service/pkg/models"
)

type UserRepo interface {
	GetUserData(ctx context.Context, id int32) (models.UserData, error)
	GetUserDataList(ctx context.Context, ids []int32) ([]models.UserData, models.NotFoundList, error)
}
