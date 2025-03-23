package main

import "fmt"

func example1() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]

	fmt.Println("Array :", array)
	fmt.Println("Slice :", slice)
}

func add() {
	table1 := []int{1, 2, 3}
	slice1 := append(table1, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("Slice added elements : ", slice1)
}

func working() {
	arr := [7]string{"This", "is", "the", "tutorial", "of", "Go", "language"}

	// Display array
	fmt.Println("Array :", arr)

	// Creating slice
	myslice := arr[1:6]

	// Display slice
	fmt.Println("Slice :", myslice)

	// Display length of the slice
	fmt.Printf("Length of the slice: %d", len(myslice))

	// Display the capacity of the sliceù
	fmt.Printf("\nCapacity of the slice : %d", cap(myslice))
}
