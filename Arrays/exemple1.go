package main

func calculateAverage(arr [6]int, size int) int {
	var sum int
	for _, value := range arr {
		sum += value
	}

	return sum / size
}

func modifyArray(arr *[5]int) {
	for i := range arr {
		arr[i]++
	}
}
