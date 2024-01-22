package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Speak()
	Move()
	Eat()
}

type Cow struct {
}

type Bird struct {
}

type Snake struct {
}

func (C Cow) Speak() {
	fmt.Println("moo")
}

func (C Cow) Eat() {
	fmt.Println("grass")
}

func (C Cow) Move() {
	fmt.Println("walk")
}

func (C Bird) Speak() {
	fmt.Println("peep")
}

func (C Bird) Eat() {
	fmt.Println("worms")
}

func (C Bird) Move() {
	fmt.Println("fly")
}

func (C Snake) Speak() {
	fmt.Println("hsss")
}

func (C Snake) Eat() {
	fmt.Println("mice")
}

func (C Snake) Move() {
	fmt.Println("slither")
}

func main() {
	animal := map[string]Animal{}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		qtype := strings.Split(scanner.Text(), " ")[0]
		if qtype == "exit" {
			fmt.Print("Exiting")
			return
		}

		name := strings.Split(scanner.Text(), " ")[1]
		tipe := strings.Split(scanner.Text(), " ")[2]

		qtype = strings.ToLower(qtype)
		name = strings.ToLower(name)
		tipe = strings.ToLower(tipe)

		if qtype == "newanimal" {
			switch tipe {
			case "cow":
				animal[name] = Cow{}
			case "snake":
				animal[name] = Snake{}
			case "bird":
				animal[name] = Bird{}
			}
		} else {
			switch tipe {
			case "eat":
				animal[name].Eat()
			case "speak":
				animal[name].Speak()
			case "move":
				animal[name].Move()

			}
		}
	}
}
