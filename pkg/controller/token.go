package controller

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTokenByGinContext(ctx *gin.Context) (string, error) {
	requestToken := ctx.GetHeader("Authorization")
	splitToken := strings.Split(requestToken, "Bearer")
	if len(splitToken) != 2 {
		return "", errors.New("not found token")
	}

	token := strings.TrimSpace(splitToken[1])

	return token, nil
}
