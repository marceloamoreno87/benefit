package token

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Credentials struct {
	Login string `json:"login"`
	Senha string `json:"senha"`
}

func GetBearerToken(url_api string, credentials Credentials) string {
	requestBody, err := json.Marshal(credentials)

	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post(url_api+"/login", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	return res.Header.Get("Authorization")
}
