package user

import (
	"errors"
	"fmt"

	"github.com/programzheng/base/config"
	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/model/user"
	"gorm.io/gorm"
)

type UserAuthRequest struct {
	Token string `json:"token"`
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
	modelUserLogin := &user.UserLogin{
		UserID: u.ID,
	}

	ul, err := modelUserLogin.First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if ul == nil {
		ul, err := modelUserLogin.Add()
		if err != nil {
			return nil, err
		}
		return ul, nil
	}

	ul, err = modelUserLogin.Update(map[string]interface{}{
		"user_id": u.ID,
	}, map[string]interface{}{
		"token": token.Token,
	})
	if err != nil {
		return nil, err
	}

	return ul, nil
}
