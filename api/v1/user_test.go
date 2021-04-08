package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"unsafe"
)

func TestLogin(t *testing.T) {

	login := make(map[string]interface{})

	login["username"] = "zwz"
	login["password"] = "123456"

	bytesData, err := json.Marshal(login)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	reader := bytes.NewReader(bytesData)
	reqType := "POST"
	url := "http://localhost:8080/user/login"

	request, err := http.NewRequest(reqType, url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
}
