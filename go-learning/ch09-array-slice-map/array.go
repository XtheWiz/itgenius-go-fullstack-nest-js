package main

import "fmt"

func SampleArray1() {
	var fruits [3]string

	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "orange"

	fmt.Println("Fruits:", fruits)

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
