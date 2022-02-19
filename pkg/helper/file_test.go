package helper

import (
	"testing"
)

func TestGetFileExtensionByContentType(t *testing.T) {
	ext := GetFileExtensionByContentType("image/jpeg")
	if ext != ".jpeg" {
		t.Error("fail")
		return
	}
	t.Log("success")
}
