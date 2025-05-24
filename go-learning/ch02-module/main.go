package main

import (
	"fmt"

	"github.com/google/uuid"
	greeting "github.com/xthewiz/samplego-itgenius/greeting"
)

func main() {
	// Print "Hello, World!" to the console
	println("Hello, Ch2-Module!")

	var id string = uuid.New().String()
	fmt.Println("Generated UUID:", id)

	greeting.SayGreet("John Doe")
	greeting.SayMeet("John Doe")
}
