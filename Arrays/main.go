package main

import "fmt"

func main() {
	arr := [4]string{"geek", "gfg", "Geeks1234", "GeeksforGeeks"}

	fmt.Println("Elements of the array: ")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}

	// Double Array
	multiArray()

	// Triple Array
	treeArray()

	// Count Array length
	countLength()

	// Example 1
	example()

	// Array affectation second option
	direct()

	// Array affectation using pointers
	usePointers()

	// Some of table
	table := [6]int{3, 4, 5, 6, 7, 8}
	result := calculateAverage(table, len(table))
	fmt.Println(result)

	// Function to modify an Array
	values := [5]int{1, 2, 3, 4, 5}
	modifyArray(&values)
	fmt.Println("Incremented array : ", values)

	// Exemple of slice using in array
	example1()

	// Add elements with slice
	add()

	// See how slice work
	working()
}

func multiArray() {
	arr := [3][3]string{{"c#", "C", "Python"}, {"Java", "Scala", "Perl"}, {"C++", "GO", "HTML"}}

	// Accessing the values of the
	// array using for loop

	fmt.Println("Elements of Array 1: ")
	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr[0]); y++ {
			fmt.Print(arr[x][y])
		}
	}

	// Creating a 2-dimensional
	// array using var keyword
	// and initializing a multi
	// -dimensional array using index

	var arr1 [2][2]int
	arr1[0][0] = 100
	arr1[0][1] = 200
	arr1[1][0] = 300
	arr1[1][1] = 400

	// Accessing the values of the array
	fmt.Println("Element of array 2")
	for i := 0; i < len(arr1); i++ {
		for q := 0; q < len(arr1[0]); q++ {
			fmt.Print(arr1[i][q])
		}
	}
}

func treeArray() {
	arr := [3][3][3]string{
		{
			{"test1", "test2", "test3"},
			{"uni1", "uni2", "uni3"},
			{"ro1", "ro2", "ro3"},
		},
		{
			{"test11", "test12", "test13"},
			{"uni11", "uni12", "uni13"},
			{"ro11", "ro12", "ro12"},
		}, {
			{"test111", "test112", "test113"},
			{"uni111", "uni112", "uni113"},
			{"ro111", "ro112", "ro113"},
		},
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			for x := 0; x < len(arr[1]); x++ {
				fmt.Print(arr[i][j][x])
			}
		}
	}
}

func countLength() {
	arr1 := [3]int{3, 9, 10}
	arr2 := [...]int{9, 7, 6, 4, 5, 3, 2, 4, 1, 4, 6, 7}
	arr3 := [3]int{9, 3, 5}

	// Finding the length of the
	// array using len method
	fmt.Println("Length of the array 1 is :", len(arr1))
	fmt.Println("Length of the array 2 is :", len(arr2))
	fmt.Println("Length of the array 3 is :", len(arr3))

	for i := 0; i < len(arr1); i++ {
		fmt.Print(arr1[i])
	}

	for j := 0; j < len(arr2); j++ {
		fmt.Print(arr2[j])
	}

	for x := 0; x < len(arr3); x++ {
		fmt.Print(arr3[x])
	}
}
