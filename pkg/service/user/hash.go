package user

import "github.com/programzheng/base/pkg/helper"

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
