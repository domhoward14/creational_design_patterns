package factory

import "fmt"

type Character interface {
	getSpeed() int
	getAttack() int
	getDefense() int
}

func PrintStats(c Character) {
	fmt.Println("Speed:", c.getSpeed())
	fmt.Println("Attack:", c.getAttack())
	fmt.Println("Defense:", c.getDefense())
}

func GetCharacter(character string) Character {
	switch character {
	case "warrior":
		return warrior{speed: 10, attack: 20, defense: 30}
	case "archer":
		return archer{speed: 15, attack: 25, defense: 20}
	case "mage":
		return mage{speed: 5, attack: 30, defense: 10}
	default:
		fmt.Printf("Unknown character: %s\n", character)
		return nil
	}
}
