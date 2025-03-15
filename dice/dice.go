package dice

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"time"
)

var possibleSides = []int{4, 6, 8, 10, 12, 20, 100}

var source = rand.NewSource(time.Now().UnixNano())
var rng = rand.New(source)

type Dice struct {
	Amount int
	Sides  int
}

func CreateDice(amount, sides int) (*Dice, error) {
	if slices.Contains(possibleSides, sides) {
		return &Dice{
			Amount: amount,
			Sides:  sides,
		}, nil
	}
	return nil, errors.New("invalid dice, please try again")
}

func (dice Dice) Roll() {
	rollSum := 0
	for i := range dice.Amount {
		roll := rng.Intn(dice.Sides) + 1
		rollSum += roll
		fmt.Printf("Dice %d rolled a: %d\n", i+1, roll)
	}
	fmt.Printf("Total: %d\n", rollSum)
}
