package test

/**
测试之前必须先启动http服务
*/

import (
	"encoding/base64"
	"testing"

	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/http"
)

var auths string

func init() {
	b := token + ":123"
	s := base64.StdEncoding.EncodeToString([]byte(b))
	auths = "Basic " + s
}

func Test_Create(t *testing.T) {

	url := host + "/blog/new"

	var data = []byte(`{
   "subject":"标题",
   "blog":"测试内容xxx",
   "author":"leking"
   }`)

	b, _ := common.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := http.RequestJSON("POST", url, data, http.Header{Key: "Authorization", Value: auths})
	if err != nil {
		t.Error(err)
	}

	s, _ := common.PrettyJSON(result)
	t.Logf("Response: %s", s)
}

func Test_Find(t *testing.T) {

	url := host + "/blog/0/10"

	var data = []byte(`{}`)
	b, _ := common.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := http.RequestJSON("GET", url, data, http.Header{Key: "Authorization", Value: auths})
	if err != nil {
		t.Error(err)
	}

	s, _ := common.PrettyJSON(result)
	t.Logf("Response: %s", s)
}

func Test_Update(t *testing.T) {

	url := host + "/blog/update/" + updateID

	var data = []byte(`{
   "subject":"标题2",
   "blog":"测试内容yyy"
   }`)

	b, _ := common.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := http.RequestJSON("PUT", url, data, http.Header{Key: "Authorization", Value: auths})
	if err != nil {
		t.Error(err)
	}

	s, _ := common.PrettyJSON(result)
	t.Logf("Response: %s", s)
}

func Test_Delete(t *testing.T) {

	url := host + "/blog/delete/" + deleteID

	var data = []byte(`{}`)
	b, _ := common.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := http.RequestJSON("DELETE", url, data, http.Header{Key: "Authorization", Value: auths})
	if err != nil {
		t.Error(err)
	}

	s, _ := common.PrettyJSON(result)
	t.Logf("Response: %s", s)
}
