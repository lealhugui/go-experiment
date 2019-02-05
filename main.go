package main

import "fmt"

func main() {
	var chosenUndead WarriorOfSunlight = Undead{"Solaire"}
	fmt.Printf("%s\n", chosenUndead)
	fmt.Printf("%s\n", chosenUndead.PraiseTheSun())
}

// WarriorOfSunlight is an Interface for the covenant of the sun
type WarriorOfSunlight interface {
	PraiseTheSun() string
}

// Undead is the common race
type Undead struct {
	name string
}

// PraiseTheSun is an action that the Undead do as they are warrriors of sunlight
func (c Undead) PraiseTheSun() string {
	return fmt.Sprint("\\[T]/ PRAISE THE SUN \\[T]/")
}

// Implementation of the String interface on the Undead struct
func (c Undead) String() string {
	return fmt.Sprintf("%s-->", c.name)
}
