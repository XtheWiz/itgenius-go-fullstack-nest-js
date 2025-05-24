package greeting

import (
	"fmt"

	personal "github.com/xthewiz/samplego-itgenius/greeting/internal"
)

func SayGreet(name string) {
	personal.PrintPersonal()
	fmt.Println("Say Hello to " + name + "! from SayGreet function")
}
