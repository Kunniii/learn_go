package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func NewAnimal(food, locomotion, noise string) *Animal {
	return &Animal{food, locomotion, noise}
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func printResult(a *Animal, info string) {
	if info == "eat" {
		a.Eat()
	}
	if info == "move" {
		a.Move()
	}
	if info == "speak" {
		a.Speak()
	}
}

func createInputString() func(prompt string) (string, string) {
	reader := bufio.NewReader(os.Stdin)
	return func(prompt string) (string, string) {
		fmt.Print(prompt)
		inp, _ := reader.ReadString('\n')
		inp = strings.TrimSuffix(inp, "\n")
		a := strings.Split(inp, " ")[0]
		b := strings.Split(inp, " ")[1]
		return strings.ToLower(a), strings.ToLower(b)
	}
}

func main() {
	cow := NewAnimal("grass", "walk", "moo")
	bird := NewAnimal("worms", "fly", "peep")
	snake := NewAnimal("mice", "slither", "hsss")

	fmt.Println("Get information of `cow`, 'bird`, and `snake`")
	fmt.Println("The information are `eat`, 'move`, and `speak`")

	input := createInputString()
	name, info := input("> ")

	if name == "cow" {
		printResult(cow, info)
	}
	if name == "snake" {
		printResult(snake, info)
	}
	if name == "bird" {
		printResult(bird, info)
	}

}
