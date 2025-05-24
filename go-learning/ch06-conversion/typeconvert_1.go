package main

import "fmt"

func TypeConvert1() {
	var i int = 42
	var f float64 = float64(i)
	fmt.Printf("i: %d, f: %.2f\n", i, f)
}
