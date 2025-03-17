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
	Name   string
	Amount int
	Sides  int
}

func (dice Dice) String() string {
	return fmt.Sprintf("%s: %d D%d(s)", dice.Name, dice.Amount, dice.Sides)
}

var DiceSets = make(map[string]Dice)

func NewDice(name string, amount, sides int) (*Dice, error) {
	if slices.Contains(possibleSides, sides) {
		if sides == 100 {
			amount = 1
		}
		return &Dice{
			Name:   name,
			Amount: amount,
			Sides:  sides,
		}, nil
	}
	return nil, errors.New("invalid dice, please try again")
}

func (dice Dice) Roll() {
	rollSum := 0
	for i := 0; i < dice.Amount; i++ {
		roll := rng.Intn(dice.Sides) + 1
		rollSum += roll
		fmt.Printf("Dice %d rolled a: %d\n", i+1, roll)
		time.Sleep(time.Duration(100 * time.Millisecond)) // TESTING SMALL DELAY BETWEEN EACH DICE ROLL
	}
	fmt.Printf("Total: %d\n", rollSum)
}

func (dice Dice) SaveDice(name string) {
	DiceSets[name] = dice
}
