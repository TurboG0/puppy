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


func Version12() string {
	return "Hello, from v1.2.0"
}