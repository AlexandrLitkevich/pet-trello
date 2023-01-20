package mock

import (
	"github.com/AlexandrLitkevich/pet-trello/graph/model"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func NewUserServiceMock() *UserServiceMock { return &UserServiceMock{} }

func (s *UserServiceMock) GetUserById(id string) *model.User {
	args := s.Called(id)
	return args.Get(0).(*model.User)
}
