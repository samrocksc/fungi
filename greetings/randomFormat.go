package greetings

import (
	"math/rand"
)

func randomFormat() string {
	formats := []string{
		"Hi %v, Welcome!",
		"Hello %v, Great to see you!",
		"Hail %v, well met!",
	}
	return formats[rand.Intn(len(formats))]
}
