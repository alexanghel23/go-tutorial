package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal interface describes the methods of an animal
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow represents a cow
type Cow struct{}

// Eat prints the food the cow eats
func (c Cow) Eat() {
	fmt.Println("grass")
}

// Move prints the locomotion method of the cow
func (c Cow) Move() {
	fmt.Println("walk")
}

// Speak prints the sound the cow makes
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird represents a bird
type Bird struct{}

// Eat prints the food the bird eats
func (b Bird) Eat() {
	fmt.Println("worms")
}

// Move prints the locomotion method of the bird
func (b Bird) Move() {
	fmt.Println("fly")
}

// Speak prints the sound the bird makes
func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake represents a snake
type Snake struct{}

// Eat prints the food the snake eats
func (s Snake) Eat() {
	fmt.Println("mice")
}

// Move prints the locomotion method of the snake
func (s Snake) Move() {
	fmt.Println("slither")
}

// Speak prints the sound the snake makes
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	// Map to store animals by name
	animals := make(map[string]Animal)

	// Main loop to handle user commands
	for {
		fmt.Print("> ")

		// Read user input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		// Split the input into command and arguments
		parts := strings.Fields(input)
		if len(parts) < 2 {
			fmt.Println("Invalid input. Please enter at least two strings.")
			continue
		}

		// Process the command based on the first string
		switch strings.ToLower(parts[0]) {
		case "newanimal":
			if len(parts) != 3 {
				fmt.Println("Invalid newanimal command. Please enter 'newanimal name type'.")
				continue
			}
			name := parts[1]
			animalType := strings.ToLower(parts[2])
			createNewAnimal(name, animalType, animals)
		case "query":
			if len(parts) != 3 {
				fmt.Println("Invalid query command. Please enter 'query name info'.")
				continue
			}
			name := parts[1]
			info := strings.ToLower(parts[2])
			queryAnimalInfo(name, info, animals)
		default:
			fmt.Println("Invalid command. Please enter 'newanimal' or 'query'.")
		}
	}
}

// createNewAnimal creates a new animal based on the provided type and adds it to the map
func createNewAnimal(name, animalType string, animals map[string]Animal) {
	var newAnimal Animal

	switch animalType {
	case "cow":
		newAnimal = Cow{}
	case "bird":
		newAnimal = Bird{}
	case "snake":
		newAnimal = Snake{}
	default:
		fmt.Println("Invalid animal type. Please enter 'cow', 'bird', or 'snake'.")
		return
	}

	animals[name] = newAnimal
	fmt.Println("Created it!")
}

// queryAnimalInfo queries the information about the animal and prints the result
func queryAnimalInfo(name, info string, animals map[string]Animal) {
	animal, exists := animals[name]
	if !exists {
		fmt.Println("Animal not found. Please create the animal using 'newanimal' command.")
		return
	}

	switch info {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Invalid information requested. Please enter 'eat', 'move', or 'speak'.")
	}
}
