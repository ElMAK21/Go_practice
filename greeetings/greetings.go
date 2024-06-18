package greeetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.

func Hello(name string) (string, error) {

	//If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// Hellos returns a map that associates each of the named people
// With a greeting message
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the hello function to get a message for each nam.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the  map, associate the retrived message with the name
		messages[name] = message
	}
	return messages, nil
}

//randomFormat returns one of a set of greetings messages, the returned
//message is selected at random

func randomFormat() string {
	//A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"hail, %v! well met!",
	}
	// Rewtuirn a randomly selected message format by specifyg
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
