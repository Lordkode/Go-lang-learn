package main

import "fmt"

var source = [5]int{10, 20, 30, 40, 50}

func example() {
	fmt.Println("Source Array: ", source)

	for i := 0; i < len(source); i++ {
		fmt.Print(source[i])
	}

	// Creating a destination array with the same size as the source array
	destination := [5]int{}
	for i := 0; i < len(source); i++ {
		destination[i] = source[i]
	}

	// Displaying of destination array
	fmt.Println("Destination Array:", destination)
}

func direct() {
	// Copying by direct assignment
	var second_destination = source

	// Displaying the result
	fmt.Println("Source: ", source)
	fmt.Println("Second destination : ", second_destination)
}

func usePointers() {
	var use_pointers *[5]int = &source

	fmt.Println("Source : ", source)
	fmt.Println("Destination array via pointer : ", *use_pointers)
}
