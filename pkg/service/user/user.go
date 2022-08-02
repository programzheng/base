package user

import (
	"github.com/programzheng/base/pkg/helper"
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

type UserLoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (ur *UserRequest) GenerateUser() (*user.User, error) {
	ur.Password = generateHashPassword(ur.Password)

	modelUser := user.User{
		UUID:     helper.CreateUuid(),
		Account:  ur.Account,
		Password: ur.Password,
		Profile: user.UserProfile{
			Name: ur.Profile.Name,
		},
	}

	user, err := modelUser.Add()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByUUID(userUUID string) (*user.User, error) {
	user, err := (&user.User{
		UUID: userUUID,
	}).First()
	if err != nil {
		return nil, err
	}

	return user, nil
}
