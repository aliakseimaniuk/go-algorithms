package main

import "fmt"

func main() {
	array := []int{2, 7, 1, 9, 5, 4, 3, 3}
	fmt.Println("Unsorted array:", array)
	insertionSort(array)
	fmt.Println("Sorted array:", array)
}
