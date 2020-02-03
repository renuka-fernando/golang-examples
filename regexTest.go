package main

import (
	"fmt"
	"regexp"
)

func main() {
	sample := "https://registry-1.do/v3.3-/cker.io"
	reg := `\/v[\d.-]*$`

	var re = regexp.MustCompile(reg)
	s := re.ReplaceAllString(sample, "")

	fmt.Println(s)

}
