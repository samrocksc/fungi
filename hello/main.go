package main

import (
	"fmt"
	"log"

	"github.com/samrocksc/fungi/greetings"
)

// This is an exercise i followed from this series:
// https://golang.org/doc/tutorial/create-module
func main() {
	// creating a prefix for the upcoming errors
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// make sure you are accounting for the error here
	// message, err := greetings.Hello("gladys")

	messages, err := greetings.Hellos(names)

	// check if there is an err at this point, and throw it if so
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
