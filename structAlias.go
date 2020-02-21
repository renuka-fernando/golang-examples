package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	alice1 := &Person{"Alice", 30}
	alice2 := *alice1
	fmt.Println(*alice1 == alice2) // => true, they have the same field values
	fmt.Println(alice1 == &alice2) // => false, they have different addresses

	alice2.Age += 10
	fmt.Println(*alice1 == alice2)
	fmt.Println(alice1.Age)
	fmt.Println(alice2.Age)
}
