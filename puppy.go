package puppy

import (
	"github.com/TurboG0/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof!"
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}


func Version10() string {
	return "Hello, from v1.0.0"
}