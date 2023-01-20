package services

import (
	"github.com/AlexandrLitkevich/pet-trello/graph/model"
)

type userService struct {
}

type UserService interface {
	GetUserById(id string) *model.User
}

var mock []*model.User = []*model.User{
	{"1", "one"},
	{"2", "fima"},
	{"3", "geor"},
	{"4", "sasha"},
	{"12", "on3e"},
	{"22", "fim3a"},
	{"33", "ge3or"},
	{"45", "sas3ha"},
}

// NewUserService Use in server.go for init
func NewUserService() *userService {
	return &userService{}
}

func (u *userService) GetUserById(id string) *model.User {
	for _, el := range mock {
		if el.ID == id {
			return el
		}
	}
	return nil
}
