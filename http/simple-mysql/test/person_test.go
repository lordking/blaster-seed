/**
测试之前必须先启动http服务
*/
package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/http"
)

func Test_Create(t *testing.T) {

	url := stHost + "/person/new"

	var data = []byte(`{
		"name":"leking",
	  "phone": "18987871818"
	}`)

	b, _ := common.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := http.RequestJSON("POST", url, data)
	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	assert.Nil(t, err, errorMessage)

	s, _ := common.PrettyJSON(result)
	t.Logf("Response: %s", s)
}

func Test_Find(t *testing.T) {

	url := stHost + "/person/leking"

	var data = []byte(`{}`)

	b, _ := common.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := http.RequestJSON("GET", url, data)
	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	assert.Nil(t, err, errorMessage)

	s, _ := common.PrettyJSON(result)
	t.Logf("Response: %s", s)

}

func Test_Update(t *testing.T) {

	url := stHost + "/person/update/leking"

	var data = []byte(`{
		"phone": "18987871111"
	}`)

	b, _ := common.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := http.RequestJSON("PUT", url, data)
	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	assert.Nil(t, err, errorMessage)

	s, _ := common.PrettyJSON(result)
	t.Logf("Response: %s", s)

}

func Test_Delete(t *testing.T) {

	url := stHost + "/person/delete/leking"

	var data = []byte(`{}`)

	b, _ := common.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := http.RequestJSON("DELETE", url, data)
	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	assert.Nil(t, err, errorMessage)

	s, _ := common.PrettyJSON(result)
	t.Logf("Response: %s", s)
}
