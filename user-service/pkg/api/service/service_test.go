package service

import "github.com/stretchr/testify/mock"

var mockRepository = new(mockRepository)

type MockRepository struct {
	mock.Mock
}
