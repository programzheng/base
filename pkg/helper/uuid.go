package helper

import "github.com/google/uuid"

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func CreateUuid() string {
	id := uuid.New()
	return id.String()
}
