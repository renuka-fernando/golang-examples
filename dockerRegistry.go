package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/heroku/docker-registry-client/registry"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:5000"
	username := "" // anonymous
	password := "" // anonymous
	hub, err := registry.New(url, username, password)
	tags, err := hub.Tags("my-repo-renuka/registry-renuka")
	fmt.Println(tags, err)

	cred, err := json.Marshal(map[string]string{
		"username": "",
		"password": "",
	})
	resp, err := http.Post("https://hub.docker.com/v2/users/login/", "application/json", bytes.NewBuffer(cred))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	jwt := &struct {
		token string
	}{}
	_ = json.Unmarshal(body, jwt)

	fmt.Println(jwt, resp.Header, resp.StatusCode, resp.Status, err)

	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://hub.docker.com/v2/repositories/{USERNAME}", nil)
	request.Header.Set("Authorization", "JWT "+jwt.token)
	res, _ := client.Do(request)
	defer res.Body.Close()
	body2, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body2))

}
