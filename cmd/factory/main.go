package main

import (
	"fmt"
	"os"

	"github.com/domhoward14/creational_design_patterns/internal/factory"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a character as a command-line argument")
		return
	}

	// Factory method right here. Allows client code to interact only with the interface which hides
	// implementation details of all the concrete objects and allows for easy expansion of futher characters.
	c := factory.GetCharacter(os.Args[1])
	factory.PrintStats(c)
}
