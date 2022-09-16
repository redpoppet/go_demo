package main

import (
	"fmt"
	"greetings"
	"hello/api"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	api.Echo_echo()
}
