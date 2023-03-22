package benefit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	msbenefit "github.com/marceloamoreno87/benefit/proto"
)

func GetBenefits(url_api string, bearerToken string, doc string) []*msbenefit.User {

	var result []*msbenefit.User
	var wg sync.WaitGroup
	user := make(chan *msbenefit.User)
	chunks := strings.Split(doc, ",")
	for _, c := range chunks {
		wg.Add(1)
		go requestDoc(url_api, bearerToken, c, user, &wg)
		result = append(result, <-user)

	}
	wg.Wait()
	return result
}

func requestDoc(url_api string, bearerToken string, c string, user chan *msbenefit.User, wg *sync.WaitGroup) {
	req, err := http.NewRequest("GET", url_api+"/offline/listagem/"+c, nil)
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
	user <- &body
	wg.Done()

}
