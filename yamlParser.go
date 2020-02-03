package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

func main() {
	yamlString := `
name: Renuka
age: 38
address:
  line1: 332
  line2: 132
`
	student := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(yamlString), &student)
	student["address"].(map[interface{}]interface{})["line1"] = "Bandaragama"

	fmt.Println(student, err)

	out, err := yaml.Marshal(student)
	fmt.Println(string(out), err)

}
