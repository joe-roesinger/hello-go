package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	message, _ := greetings.Hello("Steve")

	fmt.Println(message)
}
