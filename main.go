package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Terry-Bridge/ttrpg-assistant/dice"
)

func main() {

	// CODE BELOW IS USED SOLELY FOR TESTING PURPOSES AT THIS TIME!
	// attackDice, err := dice.NewDice("Attack", 2, 6)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Printf("Your newly created set of dice consist of %d %d sided dice\n", attackDice.Amount, attackDice.Sides)
	// }

	// percentiles, err2 := dice.NewDice("percentiles", 2, 100)
	// if err2 != nil {
	// 	fmt.Println("Error:", err2)
	// } else {
	// 	fmt.Printf("Your newly created set of dice consist of %d %d sided dice\n", percentiles.Amount, percentiles.Sides)
	// }

	// attackDice.Roll()
	// percentiles.Roll()
	// attackDice.SaveDice("Attack")
	// percentiles.SaveDice("Percentiles")

	// fmt.Println()

	// testDice, err3 := dice.NewDice(dice.DiceSets["Attack"].Name, dice.DiceSets["Percentiles"].Amount, dice.DiceSets["Attack"].Sides)
	// if err3 != nil {
	// 	fmt.Println("Error:", err3)
	// }
	// testDice.SaveDice("Test")

	// for k, v := range dice.DiceSets {
	// 	fmt.Printf("%s: %v\n", k, v)
	// }

	// fmt.Println()
	// fmt.Println(attackDice.Name)
	// fmt.Println(percentiles.Name)
	// fmt.Println("*******************")

	// fmt.Println("Name: ", testDice.Name)
	// fmt.Println("Amount:", testDice.Amount)
	// fmt.Println("Sides: ", testDice.Sides)
	// testDice.Roll()

	// player, err4 := characterCreator.CreateCharacter("Dutch", "Ranger", "Elf", "Neutral", "Some background")
	// if err4 != nil {
	// 	fmt.Println("Error:", err4)
	// }

	// fmt.Println(player)
	// fmt.Println(player.Stats)

	fmt.Println("What do you want to call this set of dice?")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	diceName := scanner.Text()

	fmt.Println("How many sides do you want your dice to have?")
	scanner.Scan()

	sides := scanner.Text()
	diceSides, err := strconv.Atoi(sides)
	if err != nil {
		fmt.Println("Please enter a valid number")
	}

	fmt.Println("How many dice?")
	scanner.Scan()

	amount := scanner.Text()
	diceAmount, err := strconv.Atoi(amount)
	if err != nil {
		fmt.Println("Please enter a valid number")
	}

	newDice, err := dice.NewDice(diceName, diceAmount, diceSides)
	if err != nil {
		fmt.Println("There was an error creating the dice")
	}

	fmt.Printf("You've created a set of dice named: %s\nIt consists of %d d%d(s)\n", newDice.Name, newDice.Amount, newDice.Sides)
	newDice.Roll()

	fmt.Println(newDice)
}
