package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)

	for {
		fmt.Print("Enter an integer (or 'X' to quit): ")
		var input string
		fmt.Scanln(&input)

		if input == "X" {
			break
		}

		integer, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}

		slice = append(slice, integer)

		sort.Ints(slice)

		fmt.Println("Sorted slice:", slice)
	}

	fmt.Println("Goodbye!")
}
