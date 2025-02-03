package main

import (
	"fmt"

	"github.com/Terry-Bridge/ttrpg-assistant/character"
)

func main() {
	fmt.Println("Welcome to my ttrpg assistant")

	//Code below is used for testing purposes, DELETE WHEN NO LONGER NEEDED
	c1, _ := character.CreateCharacter("Dutch", "RaNGer", "ELf", 1)

	fmt.Printf("%s\n%s\n%s\n%v\n", c1.Name, c1.Class, c1.Race, c1.Level)
	fmt.Println(c1.Stats)

	c1.AddAlignment("an evil dude")
	c1.AddBackground("abandoned as a baby and raised by a group of Orcs")

	fmt.Println(c1.Alignment)
	fmt.Println(c1.Background)

}
