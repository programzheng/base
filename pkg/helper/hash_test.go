package helper

import (
	"testing"
)

func TestCreateMD5(t *testing.T) {
	hash := CreateMD5("test")
	if hash != "098f6bcd4621d373cade4e832627b4f6" {
		t.Error("fail")
		return
	}
	t.Log("success")
}
