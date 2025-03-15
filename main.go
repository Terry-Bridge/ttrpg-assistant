package main

import (
	"fmt"

	"github.com/Terry-Bridge/ttrpg-assistant/dice"
)

func main() {

	// CODE BELOW IS USED SOLELY FOR TESTING PURPOSES AT THIS TIME!
	attackDice, err := dice.CreateDice(4, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Your newly created set of dice consist of %d %d sided dice\n", attackDice.Amount, attackDice.Sides)
	}

	attackDice.Roll()

}
