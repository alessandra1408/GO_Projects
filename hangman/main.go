package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectLanguage() {
	//usar switch

}

func getValidWords() []string {

	//se for array pt-br
	//fazer um loop for para:
	//todas as posições do array menores que 3 serem excluídas
	//gerar um novo array com essas novas palavras

	var arrayInput = []string{}
	var validWords = []string{}
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

	for i := 0; i < len(arrayInput); i++ {
		if len(arrayInput[i]) >= 3 {
			validWords = append(validWords, arrayInput[i])
		}
	}

	return validWords
}

func hangman(validArray []string) {

	rand.Seed(time.Now().UTC().UnixNano())
	randWord := rand.Intn(len(validArray))
	fmt.Println(validArray[randWord])
}

func main() {
	selectLanguage()
	validArray := getValidWords()
	hangman(validArray)
}
