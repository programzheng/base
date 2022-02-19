package helper

import (
	"testing"
)

func TestCreateUuid(t *testing.T) {
	for i := 0; i < 3; i++ {
		idString := CreateUuid()
		if !IsValidUUID(idString) {
			t.Errorf("fail: %v", idString)
		}
	}
	t.Log("success")
}
