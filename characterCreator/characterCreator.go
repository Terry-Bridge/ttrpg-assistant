package characterCreator

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

var AvailableRaces = []string{"aasimar", "dragonborn", "dwarf", "elf", "gnome", "goliath", "halfling", "human", "orc", "tiefling"}
var AvailableClasses = []string{"barbarian", "bard", "cleric", "druid", "fighter", "monk", "paladin", "ranger", "rogue", "sorcerer", "warlock", "wizard"}
var AvailableAlignments = []string{"lawful good", "neutral good", "chaotic good", "lawful neutral", "neutral", "chaotic neutral", "lawful evil", "neutral evil", "chaotic evil"}

type Stats struct {
	Strength     uint
	Dexterity    uint
	Constitution uint
	Intelligence uint
	Wisdom       uint
	Charisma     uint
}

func (stats Stats) String() string {
	return fmt.Sprintf("Stats:\n\tStrength: %d\n\tDexterity: %d\n\tConstitution: %d\n\tIntelligence: %d\n\tWisdom: %d\n\tCharisma: %d\n", stats.Strength, stats.Dexterity, stats.Constitution, stats.Intelligence, stats.Wisdom, stats.Charisma)
}

type Character struct {
	Name       string
	Class      string
	Level      uint
	Background string
	Race       string
	Alignment  string
	Experience int
	Stats      Stats
}

func (character Character) String() string {
	return fmt.Sprintf("%s: A level %d %s %s", character.Name, character.Level, character.Race, character.Class)
}

func CreateCharacter(name, class, race, alignment, background string) (*Character, error) {
	if slices.Contains(AvailableClasses, strings.ToLower(class)) && slices.Contains(AvailableRaces, strings.ToLower(race)) {
		return &Character{
			Name:       name,
			Class:      class,
			Race:       race,
			Alignment:  alignment,
			Background: background,
			Level:      1,
			Experience: 0,
			Stats: Stats{
				Strength:     10,
				Dexterity:    10,
				Constitution: 10,
				Intelligence: 10,
				Wisdom:       10,
				Charisma:     10,
			},
		}, nil
	}
	return nil, errors.New("error creating character. please choose a valid class and race")

}
