package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func welcomePlayer() {
	asciiArt := `
	 _   _    _    _   _  ____ __  __    _    _   _
	| | | |  / \  | \ | |/ ___|  \/  |  / \  | \ | |
	| |_| | / _ \ |  \| | |  _| |\/| | / _ \ |  \| |
	|  _  |/ ___ \| |\  | |_| | |  | |/ ___ \| |\  |
	|_| |_/_/   \_\_| \_|\____|_|  |_/_/   \_\_| \_|
	`
	fmt.Println(asciiArt)
}

func selectLanguage() []string {
	var arrayInput = []string{}
	languageOption := 0

	fmt.Println("\n** Language Selection **")
	fmt.Println("1 - English")
	fmt.Println("2 - Português")
	languageOption, err := fmt.Scan(&languageOption)

	if err != nil {
		fmt.Println("elect only one of the options.")
	}

	switch languageOption {
	case 1:
		arrayInput = wordsEN
	case 2:
		arrayInput = wordsPT
	}

	return arrayInput

}

func getValidWords(arrayInput []string) []string {

	var validWords []string

	for i := 0; i < len(arrayInput); i++ {
		if len(arrayInput[i]) >= 3 {
			validWords = append(validWords, arrayInput[i])
		}
	}

	return validWords

}

func generatWord(validArray []string) string {

	rand.Seed(time.Now().UTC().UnixNano())
	randWord := rand.Intn(len(validArray))
	generatedWord := validArray[randWord]
	generatedWord = string(generatedWord)

	return generatedWord
}

func endGame(result bool) {
	fmt.Println("ending the game...")
	time.Sleep(2 * time.Second)
	fmt.Print("\033[H\033[2J")

	if result {
		asciiArt := "Congrats. You Won!"
		fmt.Println(asciiArt)
	} else {
		fmt.Println("You lose!")
	}
}

func game(validArray []string) bool {
	var letters string
	var usedLetters []string
	generatedWord := generatWord(validArray)
	life := 5
	match := false

	wordShow := make([]string, 0)

	for i := 0; i < len(generatedWord); i++ {
		wordShow = append(wordShow, "-")
	}

	for {
		hit := false
		fmt.Println(wordShow)
		fmt.Scanln(&letters)
		letters = strings.ToLower(letters)
		usedLetters = append(usedLetters, letters)
		usedLettersConcat := strings.Join(usedLetters, " ")
		fmt.Print("\033[H\033[2J")
		fmt.Printf("\n==============\nLIFES: %d\nletters used: %v\n==============\n", life, usedLettersConcat)
		fmt.Println()
		fmt.Println("Try to guess de word!")

		//verificadores
		if len(letters) != 1 {
			fmt.Println("You can only use one letter at a time!")
			continue
		}

		for k, v := range generatedWord {
			if string(v) == letters {
				wordShow[k] = string(v)
				hit = true
			}
		}

		if !hit {
			life--
		}

		if life <= 0 {
			match = false
			break
		}

		count := 0
		for _, v := range wordShow {
			if v != "-" {
				count++

				if count == len(wordShow) {
					match = true
					break
				}

			} else {
				count = 0
			}
		}

		if match {
			break
		}
	}

	return match
}

func main() {
	welcomePlayer()
	arrayInput := selectLanguage()
	validArray := getValidWords(arrayInput)
	result := game(validArray)
	endGame(result)
}
