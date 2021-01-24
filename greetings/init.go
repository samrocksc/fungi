package greetings

import (
	"math/rand"
	"time"
)

// Init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}
