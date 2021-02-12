package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, erro := greetings.Hello("")

	if erro != nil {
		log.Println(erro)
	}

	fmt.Println(message)
}
