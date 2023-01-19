package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// var generated_number int
var maxNumber int
var feedback string

func computerGuess() {
	//create a var to get the rand
	//create two var to set the minNumber and maxNumber number
	//create a var to get feedback from user
	//create a loop to verificate that feedback is equal to computer rand number
	minNumber := 1

	fmt.Println("Do you want to I guess between 1 to...? ")
	fmt.Scanln(&maxNumber)

	generatedNumber := rand.Intn(maxNumber)

	for strings.ToLower(feedback) != "c" {
		if maxNumber != minNumber {
			rand.Seed(time.Now().UTC().UnixNano())
			generatedNumber = rand.Intn(maxNumber-minNumber) + minNumber

			fmt.Printf("Is %d too high (H), too low (L) or correct (c)? ", generatedNumber)
			fmt.Scanln(&feedback)

			if strings.ToLower(feedback) == "h" {
				maxNumber = generatedNumber - 1
			} else if strings.ToLower(feedback) == "l" {
				minNumber = generatedNumber + 1
			}
		} else {
			generatedNumber = minNumber
			feedback = "c"
			fmt.Printf("\nYay! The computer guessed your number, %d, correctly! :)", generatedNumber)
		}

	}

}

func main() {
	computerGuess()
}
