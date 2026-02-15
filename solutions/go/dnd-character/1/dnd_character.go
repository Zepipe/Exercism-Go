package dndcharacter

import (
    "math"
	"math/rand"
	"sort"
	"time"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

func Modifier(score int) int {
    return int(math.Floor(float64(score-10) / 2.0))
}

func Ability() int {
	rolls := make([]int, 4)
	for i := 0; i < 4; i++ {
		rolls[i] = rand.Intn(6) + 1
	}
	
	sort.Ints(rolls)
	
	// Sum the three largest (indices 1, 2, 3 after sorting)
	return rolls[1] + rolls[2] + rolls[3]
}

func GenerateCharacter() Character {
	rand.Seed(time.Now().UnixNano())
	
	c := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	
	c.Hitpoints = 10 + Modifier(c.Constitution)
	
	return c
}