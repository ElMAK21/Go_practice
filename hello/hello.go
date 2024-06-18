package main

import (
	"fmt"

	greeetings "example.com/greetings"

	"log"
)

func main() {
	//Set properties of the predefined Logger, including
	//the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greeetings: ")
	log.SetFlags(0)
	// A slice of names.
	names := []string{"Kevin", "Jorge", "Joosue"}

	// request greeting messages for the names.

	// Get a greeting message and print it.

	messages, err := greeetings.Hellos(names)
	// if an error was returned print it to the console and
	// exit the program

	if err != nil {
		log.Fatal(err)
	}
	//if no error was returned, printy the returned map of messages
	// to the console

	fmt.Println(messages)
}
