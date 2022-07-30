package user

import (
	"fmt"

	"github.com/programzheng/base/config"
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

func generateHashPassword(password string) string {
	hash := helper.CreateHash(password)

	return hash
}

func checkHashPassword(hash string, password string) error {
	err := helper.CheckHash(hash, password)
	if err != nil {
		return err
	}

	return err
}

func (ulgr *UserLoginRequest) Login() (*helper.Token, error) {
	modelUser := user.User{
		Account: ulgr.Account,
	}

	u, err := modelUser.First()
	if err != nil {
		return nil, err
	}

	//check hash password
	err = checkHashPassword(u.Password, ulgr.Password)
	if err != nil {
		return nil, fmt.Errorf("check password is error: %v", err)
	}

	//generate jwt token
	secret := []byte(config.Cfg.GetString("JWT_SECRET"))
	expiresSeconds := helper.ConvertToInt64(config.Cfg.GetString("JWT_EXPIRES_SECONDS"))
	token := helper.CreateJWT(secret, expiresSeconds)

	//add user login record
	_, err = loginRecord(u, token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func loginRecord(u *user.User, token *helper.Token) (*user.UserLogin, error) {
	var modelUserLogin user.UserLogin
	_, err := modelUserLogin.BatchForceDelete(map[string]interface{}{
		"user_id": u.ID,
	})
	if err != nil {
		return nil, err
	}

	modelUserLogin.UserID = u.ID
	modelUserLogin.Token = token.Token

	ul, err := modelUserLogin.Add()

	if err != nil {
		return nil, err
	}

	return ul, nil
}
