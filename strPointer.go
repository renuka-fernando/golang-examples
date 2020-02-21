package main

import "fmt"

func main() {
	var strPointer = new(string)
	*strPointer = "hello"

	fmt.Println(*strPointer)
}
