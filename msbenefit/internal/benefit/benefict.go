package benefit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	msbenefit "github.com/marceloamoreno87/msbenefit/proto"
)

func GetBenefits(url_api string, bearerToken string, doc string) msbenefit.User {
	req, err := http.NewRequest("GET", url_api+"/offline/listagem/"+doc, nil)
	req.Header.Add("Authorization", bearerToken)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	body := msbenefit.User{}
	json.Unmarshal([]byte(respBody), &body)
	return body
}
