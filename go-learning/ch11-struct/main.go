package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() string {
	return "Hello! " + p.Name
}

func (p *Person) HaveBirthday() {
	p.Age += 1
}

func main() {
	p := Person{
		Name: "John",
		Age:  30,
	}

	fmt.Printf("Person: %v\n", p)

	fmt.Println("Greeting:", p.SayHello())

	fmt.Println("HaveBirthday before:", p.Age)
	p.HaveBirthday()
	fmt.Println("HaveBirthday after:", p.Age)
}
