package user

import (
	"github.com/programzheng/base/pkg/model/user"
)

type UserRequest struct {
	Account  string      `json:"account"`
	Password string      `json:"password"`
	Profile  UserProfile `json:"profile"`
}

type UserProfile struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

func (uq *UserRequest) GenerateUser() (*user.User, error) {
	modelUser := user.User{
		Account:  uq.Account,
		Password: uq.Password,
		Profile: user.UserProfile{
			Name: uq.Profile.Name,
		},
	}

	user, err := modelUser.Add()
	if err != nil {
		return nil, err
	}

	return user, nil
}
