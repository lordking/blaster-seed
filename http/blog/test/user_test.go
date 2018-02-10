package test

import (
	"testing"

	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/http"
)

func Test_Login(t *testing.T) {

	url := host + "/user/login"

	data := []byte(`{
		"username": "admin",
		"password": "admin"
	}`)

	b, _ := common.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := http.RequestJSON("POST", url, data)
	if err != nil {
		t.Error(err)
	}

	s, _ := common.PrettyJSON(result)
	t.Logf("Response: %s", s)
}
