package helper

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/gob"
	"fmt"

	log "github.com/sirupsen/logrus"

	"golang.org/x/crypto/bcrypt"
)

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func CreateSHA1(secret interface{}) string {
	// 產生模式
	hash := sha1.New()

	// 轉換字串
	hash.Write([]byte(secret.(string)))

	// 最終hash結果
	bs := hash.Sum(nil)

	//將byte轉為16進制
	result := fmt.Sprintf("%x", bs)
	return result
}

func CreateMD5(secret interface{}) string {
	secret = ConvertToString(secret)
	// 產生模式
	hash := md5.New()

	// 轉換字串
	hash.Write([]byte(secret.(string)))

	// 最終hash結果
	bs := hash.Sum(nil)

	//將byte轉為16進制
	result := fmt.Sprintf("%x", bs)
	return result
}

func CreateHash(secret string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

func CheckHash(hash string, secret string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	return err
}
