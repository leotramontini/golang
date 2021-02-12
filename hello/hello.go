package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	// // Get a greeting message and print it.
	// message, erro := greetings.Hello("Dylan")

	// if erro != nil {
	// 	log.Println(erro)
	// }

	// fmt.Println(message)

	names := []string{"Gladys", "Samantha", "Darrin"}
	message, erro := greetings.Hellos(names)

	if erro != nil {
		log.Println(erro)
	}

	fmt.Println(message)
}
