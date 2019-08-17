package function

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var (
	secret string
)

type Token struct {
	Token string `json:"token"`
}

func init() {
	secret = viper.Get("JWT_SECRET").(string)
}

func CreateJWT() (token Token) {
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	jwtToken.Claims = claims

	jwtTokenString, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}

	token.Token = jwtTokenString

	return token
}
