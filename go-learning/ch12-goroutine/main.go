package main

import "time"

func brewCoffee() {
	println("Brewing coffee...")
	time.Sleep(2 * time.Second)
	println("Coffee is ready!")
}

func bakeBread() {
	println("Baking bread...")
	time.Sleep(3 * time.Second)
	println("Bread is ready!")
}

func washDishes() {
	println("Washing dishes...")
	time.Sleep(1 * time.Second)
	println("Dishes are clean!")
}

func main() {
	println("Starting the coffee and bread preparation...")

	// Start brewing coffee in a goroutine
	go brewCoffee()

	// Start baking bread in a goroutine
	go bakeBread()

	// Start washing dishes in a goroutine
	go washDishes()

	// Wait for a while to let the goroutines finish
	time.Sleep(5 * time.Second)

	println("All done!")
}
