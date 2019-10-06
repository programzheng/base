package function

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var (
	secret []byte
)

type Token struct {
	Token string
	Exp   int64
}

func init() {
	secret = []byte(viper.Get("JWT_SECRET").(string))
}

func CreateJWT() (token Token) {
	exp := time.Now().Add(time.Hour * time.Duration(1)).Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"test": "test",
		"exp":  exp,
		"iat":  time.Now().Unix(),
		"nbf":  time.Now().Unix(),
	})
	jwtTokenString, err := jwtToken.SignedString(secret)
	if err != nil {
		log.Fatal(err)
	}

	token.Token = jwtTokenString
	token.Exp = exp

	return token
}

func ValidJSONWebToken(requestToken string) (value interface{}, err error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		// validate the alg
		if _, err := token.Method.(*jwt.SigningMethodHMAC); !err {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if claims, err := token.Claims.(jwt.MapClaims); err && token.Valid {
		value = claims
	}

	return

}
