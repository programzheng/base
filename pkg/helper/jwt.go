package helper

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/golang-jwt/jwt"
)

var (
	secret []byte
)

type Token struct {
	Token string `json:"token"`
	Exp   int64  `json:"exp"`
}

func CreateJWT(sercrt []byte, exp int64) (token Token) {
	if exp < 0 {
		exp = 0
	}
	exp = time.Now().Add(time.Second * time.Duration(exp)).Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": exp,
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	})
	jwtTokenString, err := jwtToken.SignedString(secret)
	if err != nil {
		log.Fatal(err)
	}

	token.Token = jwtTokenString
	token.Exp = exp

	return token
}

func ValidJSONWebToken(requestToken string) bool {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		// validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
	if err != nil {
		return false
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return ok
	} else {
		return false
	}

}
