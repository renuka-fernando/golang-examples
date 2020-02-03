package main

import (
	"fmt"
	"github.com/heroku/docker-registry-client/registry"
)

func main() {
	url := "http://localhost:5000"
	username := "" // anonymous
	password := "" // anonymous
	hub, err := registry.New(url, username, password)
	tags, err := hub.Tags("my-repo-renuka/registry-renuka")
	fmt.Println(tags, err)
}
