package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Lee")
	fmt.Println(message)
}
