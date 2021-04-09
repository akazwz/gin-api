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

	reqType := "POST"
	url := "http://localhost:8000/user/login"

	resp := RequestJsonHelper(reqType, url, login)

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	str := (*string)(unsafe.Pointer(&respBytes))
	got := *str
	want := `{"code":0,"data":{"username":"zwz","password":"123456"},"msg":"Login Success"}`

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func TestRegister(t *testing.T) {
	register := make(map[string]interface{})

	register["username"] = "oreo"
	register["password"] = "123456"

	reqType := "POST"
	url := "http://localhost:8000/user/register"

	_ = RequestJsonHelper(reqType, url, register)

	url = "http://localhost:8000/user/login"

	resp := RequestJsonHelper(reqType, url, register)

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	str := (*string)(unsafe.Pointer(&respBytes))
	got := *str
	want := `{"code":0,"data":{"username":"oreo","password":"123456"},"msg":"Login Success"}`

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}

}

func RequestJsonHelper(reqType string, url string, jsonMap map[string]interface{}) *http.Response {
	bytesData, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	reader := bytes.NewReader(bytesData)

	request, err := http.NewRequest(reqType, url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return resp
}
