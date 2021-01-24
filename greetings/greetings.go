package greetings

import (
	"errors"
	"fmt"
)

// Hello returns the name of the greeted person in the param
func Hello(name string) (string, error) {
	// check to see if there is a string, or throw an error
	if name == "" {
		return "", errors.New("Missing a name on this string")
	}
	// create the message that we will use for the greeting
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}
