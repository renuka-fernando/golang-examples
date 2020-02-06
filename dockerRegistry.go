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
	fmt.Println(string(body), resp.Header, resp.StatusCode, resp.Status, err)
}
