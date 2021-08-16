package driver

import (
	"fmt"
	"io"
	"testing"
)

func TestGet(t *testing.T) {
	response := Get("https://localhost")
	body, _ := io.ReadAll(response.Body)

	fmt.Printf("response:%v\n", string(body))
	t.Log(response)

}
