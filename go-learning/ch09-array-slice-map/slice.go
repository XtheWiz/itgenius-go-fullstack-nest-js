package main

import "fmt"

func SampleSlice1() {
	var fruits = []string{"apple", "banana", "cherry"}
	fruits = append(fruits, "durian")
	fmt.Println(fruits)

	var numbers = []int{1, 2, 3}
	numbers = append(numbers, 4)
	numbers = append(numbers, 5)
	fmt.Println(numbers)

	subSlice := numbers[1:3]
	fmt.Println(subSlice)
}
