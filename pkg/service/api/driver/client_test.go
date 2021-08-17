package driver

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGet(t *testing.T) {
	response := Get("https://data.ntpc.gov.tw/api/datasets/308DCD75-6434-45BC-A95F-584DA4FED251/json/preview", nil)
	defer response.Body.Close()
	t.Log(response)
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Printf("response:%v\n", body)
	t.Log(string(body))

}
