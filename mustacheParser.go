package main

import (
	"fmt"
	"github.com/cbroglie/mustache"
)

func main() {
	template := `{
	"name": "{{name}}",
	"age" : {{age}}
}`

	output, _ := mustache.Render(template, map[string]string{
		"name": "renuka",
		"age":  "26",
	})

	fmt.Println(output)
}
