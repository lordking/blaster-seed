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

func Test_Hello(t *testing.T) {

	url := stHost + "/welcome/hello"

	var data = []byte(`{
		"name":"leking",
	   "content": {
	   		"Ye":"You are welcome"
	   }
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
