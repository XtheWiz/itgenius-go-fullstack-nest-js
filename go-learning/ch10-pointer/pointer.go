package main

import "fmt"

func SamplePointer1() {
	value := 10

	pointer := &value
	fmt.Printf("value: %v\n", value)
	fmt.Printf("pointer: %v\n", pointer)

	value = 20
	fmt.Println("change value to 20")
	fmt.Printf("pointer value: %v\n", *pointer)

	slice := []int{1, 2, 3}
	pointerToSlice := &slice
	fmt.Printf("slice address: %v\n", pointerToSlice)
	fmt.Printf("slice value: %v\n", *pointerToSlice)

	age := 10
	fmt.Printf("age before changed: %v\n", age)
	changeAge(&age)
	fmt.Printf("age before changed: %v\n", age)
}

func changeAge(age *int) {
	*age += 1
}

type Room struct {
	RoomNumber int
	BedColor   string
}

func changeBedColor(roomKey *Room, newColor string) {
	roomKey.BedColor = newColor
}

func SamplePointer2() {
	room101 := Room{
		RoomNumber: 101,
		BedColor:   "yellow",
	}

	fmt.Printf("before change color: %v\n", room101)
	changeBedColor(&room101, "red")
	fmt.Printf("after change color: %v\n", room101)
}
