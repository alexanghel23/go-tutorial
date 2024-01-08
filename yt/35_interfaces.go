package main

import (
	"fmt"
)

var pl = fmt.Println

// Interfaces are contracts that says if anything inherits it they must implement some defined methods

type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	pl("cat attacks its prey")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	pl("Cat says Hisss")
}

func (c Cat) HappySound() {
	pl("Cat says Purrr")
}

func main() {

	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()

	var kitty2 Cat = kitty.(Cat)

	kitty2.Attack()

	pl("Cats name:", kitty2.Name())

}
