package character

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Character struct {
	Name       string
	Class      string
	Race       string
	Level      uint
	Experience int
	Background string
	Alignment  string
	Stats      map[string]uint
}

type Race struct {
	Name    string
	Bonuses map[string]uint
}

var DefaultStats = []int{15, 14, 13, 12, 10, 8}

var AvailableRaces = map[string]Race{
	"dwarf":      {"dwarf", map[string]uint{"Constitution": 2, "Strength": 1}},
	"elf":        {"elf", map[string]uint{"Dexterity": 2, "Wisdom": 1}},
	"halfling":   {"halfling", map[string]uint{"Dexterity": 2, "Charisma": 1}},
	"human":      {"human", map[string]uint{"Strength": 1, "Dexterity": 1, "Constitution": 1}},
	"aasimar":    {"aasimar", map[string]uint{"Charisma": 2, "Wisdom": 1}},
	"dragonborn": {"dragonborn", map[string]uint{"Strength": 2, "Charisma": 1}},
	"gnome":      {"gnome", map[string]uint{"Intelligence": 2, "Dexterity": 1}},
	"goliath":    {"goliath", map[string]uint{"Strength": 2, "Constitution": 1}},
	"orc":        {"orc", map[string]uint{"Strength": 2, "Constitution": 1}},
	"tiefling":   {"tiefling", map[string]uint{"Charisma": 2, "Intelligence": 1}},
}

var AvailableClasses = map[string]struct{}{
	"artificer": {},
	"barbarian": {},
	"bard":      {},
	"cleric":    {},
	"druid":     {},
	"fighter":   {},
	"monk":      {},
	"paladin":   {},
	"ranger":    {},
	"rogue":     {},
	"sorcerer":  {},
	"warlock":   {},
	"wizard":    {},
}

// Creates a new playable charcter with placeholder stats temporarily, until they're assigned by the player.
func CreateCharacter(name, class, race string, level uint) (*Character, error) {

	lowercaseRace := strings.ToLower(race)
	lowercaseClass := strings.ToLower(class)

	_, raceExists := AvailableRaces[lowercaseRace]
	if !raceExists {
		return nil, fmt.Errorf("invalid race: %s", race)
	}

	_, classExists := AvailableClasses[lowercaseClass]
	if !classExists {
		return nil, fmt.Errorf("invalid class: %s", class)
	}

	//This is needed to return the character's attributes in title case
	title := cases.Title(language.AmericanEnglish)

	return &Character{
		Name:  name,
		Class: title.String(class),
		Race:  title.String(race),
		Level: level,
		Stats: map[string]uint{
			"Strength":     8,
			"Dexterity":    8,
			"Constitution": 8,
			"Intelligence": 8,
			"Wisdom":       8,
			"Charisma":     8,
		},
	}, nil
}

func (c *Character) AddBackground(background string) {
	c.Background = background
}

func (c *Character) AddAlignment(alignment string) {
	c.Alignment = alignment
}

//TODO: IMPLEMENT A CLASS STRUCT
//TODO: CREATE A FUNCTION TO ASSIGN DEFAULT VALUES AND ASSIGN THEM TO CHOSEN STATS, OR ALLOW THE USER TO MANUALLY ASSIGN POINTS TO THEIR SKILLS
//TODO: CREATE A DICE PACKAGE
//TODO: WITHIN THE DICE PACKAGE ALLOW THE USER TO ROLL FOR STATS (ROLL 4 D6s, REMOVE THE LOWEST DICE (REROLL 1'S) AND SUM THE OTHER 3 DICE ROLLS)
//TODO EXPAND UPON THE DICE PACKAGE TO DO THE FOLLOWING:
//	- CREATE A DICE STRUCT
//	- CREATE FUNCTIONS/METHODS
//		- ROLL
//		- TBD
