package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var filename string

	fmt.Print("Enter the name of the text file: ")
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	names := make([]name, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}
		firstName := parts[0]
		lastName := parts[1]

		n := name{firstName, lastName}

		names = append(names, n)
	}

	for _, n := range names {
		fmt.Println(n.fname, n.lname)
	}
}
