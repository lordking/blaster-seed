package benchmark

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/http"
	"github.com/lordking/blaster/log"
)

//RequestCreate hello接口的测试案例
func (t *TestCase) RequestCreate() {

	url := t.BaseURL + "/person/new"

	str := fmt.Sprintf(`{
					"name":"leking%d",
					"phone":"189aaaa%d"
				}`, t.Count, t.TesterNO)

	data := []byte(str)
	b, _ := common.PrettyJSON(data)
	log.Debugf("Request: %s", b)

	result, err := http.RequestJSON("POST", url, data)
	if err != nil {
		log.Errorf("Error: %s", err.Error())
	}

	s, _ := common.PrettyJSON(result)
	log.Debugf("Response: %s", s)
}

//RequestFind hello接口的测试案例
func (t *TestCase) RequestFind() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	url := fmt.Sprintf("%s/person/leking%d", t.BaseURL, r.Intn(t.RandLimit))

	data := []byte(`{}`)
	b, _ := common.PrettyJSON(data)
	log.Debugf("Request: %s", b)

	result, err := http.RequestJSON("GET", url, data)
	if err != nil {
		log.Errorf("Error: %s", err.Error())
	}

	s, _ := common.PrettyJSON(result)
	log.Debugf("Response: %s", s)

}

//RequestUpdate hello接口的测试案例
func (t *TestCase) RequestUpdate() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rn := r.Intn(t.RandLimit)
	url := fmt.Sprintf("%s/person/update/leking%d", t.BaseURL, rn)

	str := fmt.Sprintf(`{
					"phone":"189bbbb%d"
				}`, rn)

	data := []byte(str)
	b, _ := common.PrettyJSON(data)
	log.Debugf("Request: %s", b)

	result, err := http.RequestJSON("PUT", url, data)
	if err != nil {
		log.Errorf("Error: %s", err.Error())
	}

	s, _ := common.PrettyJSON(result)
	log.Debugf("Response: %s", s)

}

//RequestDelete hello接口的测试案例
func (t *TestCase) RequestDelete() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rn := r.Intn(t.RandLimit)
	url := fmt.Sprintf("%s/person/delete/leking%d", t.BaseURL, rn)

	data := []byte(`{}`)
	b, _ := common.PrettyJSON(data)
	log.Debugf("Request: %s", b)

	result, err := http.RequestJSON("DELETE", url, data)
	if err != nil {
		log.Errorf("Error: %s", err.Error())
	}

	s, _ := common.PrettyJSON(result)
	log.Debugf("Response: %s", s)

}
