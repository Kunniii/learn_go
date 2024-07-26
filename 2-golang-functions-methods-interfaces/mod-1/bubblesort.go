package main

import (
	"fmt"
	"strconv"
)

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
}

func main() {
	var numbers []int

	for {
		fmt.Print("Enter a number (or 'x' to exit): ")
		var input string
		fmt.Scanln(&input)

		if input == "x" {
			break
		}

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		numbers = append(numbers, number)
	}

	fmt.Println("\nYour array:", numbers)

	bubbleSort(numbers)

	fmt.Println("Sorted array:", numbers)
}
