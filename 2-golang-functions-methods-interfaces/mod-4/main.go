package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal interface with Eat, Move, and Speak methods
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow, Bird, Snake types implementing Animal interface
type (
	Cow   struct{}
	Bird  struct{}
	Snake struct{}
)

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	animals := make(map[string]Animal)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		commands := strings.Split(input, " ")

		if len(commands) != 3 {
			fmt.Println(
				"Invalid input. Please use the format: newanimal <name> <type> or query <name> <action>",
			)
			continue
		}

		switch commands[0] {
		case "newanimal":
			name, animalType := commands[1], commands[2]
			switch animalType {
			case "cow":
				animals[name] = Cow{}
			case "bird":
				animals[name] = Bird{}
			case "snake":
				animals[name] = Snake{}
			default:
				fmt.Println("Unknown animal type. Valid types are: cow, bird, snake.")
				continue
			}
			fmt.Println("Created it!")
		case "query":
			name, action := commands[1], commands[2]
			animal, exists := animals[name]
			if !exists {
				fmt.Println("No such animal exists.")
				continue
			}
			switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Unknown action. Valid actions are: eat, move, speak.")
			}
		default:
			fmt.Println(
				"Invalid command. Use 'newanimal' to create an animal or 'query' to get information.",
			)
		}
	}
}
