package main

import "fmt"

func SampleMap1() {
	var scores = make(map[string]int)
	scores["Alice"] = 90
	scores["Bob"] = 85
	scores["Charlie"] = 95

	fmt.Printf("Scores: %v\n", scores)

	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}

	grades := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 95,
		"David":   80,
	}
	fmt.Println("Grades:", grades)
}
