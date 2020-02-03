package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	mustacheGistUrl := `https://gist.githubusercontent.com/renuka-fernando/6d6c64c786e6d13742e802534de3da4e/raw/3f95ac95ae3a8ded8b40e63676e84124f9dae30c/controller_conf.yaml`

	res, err := http.Get(mustacheGistUrl)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
