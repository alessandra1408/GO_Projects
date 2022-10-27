package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func selectLanguage() []string {
	var arrayInput = []string{}
	languageOption := 0

	fmt.Println("** Language Selection **")
	fmt.Println("1 - English")
	fmt.Println("2 - Português")
	fmt.Scan(&languageOption)

	switch languageOption {
	case 1:
		arrayInput = wordsEN
	case 2:
		arrayInput = wordsPT
	}

	return arrayInput

}

func getValidWords(arrayInput []string) []string {

	//se for array pt-br
	//fazer um loop for para:
	//todas as posições do array menores que 3 serem excluídas
	//gerar um novo array com essas novas palavras

	var validWords []string

	for i := 0; i < len(arrayInput); i++ {
		if len(arrayInput[i]) >= 3 {
			validWords = append(validWords, arrayInput[i])
		}
	}

	return validWords

}

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

func hangman(validArray []string) {

	var usedLetters = []string{}

	rand.Seed(time.Now().UTC().UnixNano())
	randWord := rand.Intn(len(validArray))
	generatedWord := validArray[randWord]

	var letters string
	fmt.Println("Try to guess de word!")

	for vida := 5; vida > 0; vida-- {
		fmt.Scanln(&letters)

		if len(letters) != 1 {
			fmt.Println("You can only use one letter at a time! You lsot a life.")

		} else {

			for i := 0; i != len(generatedWord); i++ {

				if string(letters) == string(generatedWord[i]) {
					fmt.Printf("a letra %v tem na palavra!\n", letters)
				}

				fmt.Printf("_ ")
			}

			fmt.Println()
			fmt.Println(generatedWord)

			usedLetters = append(usedLetters, letters)
			usedLettersConcat := strings.Join(usedLetters, " ")
			fmt.Printf("You have used these letters: %v\n", usedLettersConcat)
		}
	}

}

func main() {
	welcomePlayer()
	arrayInput := selectLanguage()
	validArray := getValidWords(arrayInput)
	hangman(validArray)
}
