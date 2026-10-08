package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	message := greetings.Hello("Steve")
	fmt.Println(message)
}
