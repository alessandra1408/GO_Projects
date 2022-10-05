package main

import (
	"fmt"
	"math/rand"
	"time"
)

//usuario insere a opcao
//converter opcao do user no nome das opcoes
//verificar se a opcao do user foi a mesma do pc
//criar funcao de verificacao se ganhou o game

func option(player string) string {
	computer_option := ""

	if player == "r" {
		computer_option = "rock"
	} else if player == "p" {
		computer_option = "paper"
	} else {
		computer_option = "scissors"
	}
	return computer_option
}

func verification(user string, computerRand string) {
	computerOption := option(computerRand)
	userChoice := option(user)
	if user == "r" || user == "p" || user == "s" {
		if user == computerRand {
			fmt.Println("*** Its a Tie. ***")
		} else if win(user, computerRand) {
			fmt.Printf("You: %v\nComputer: %v.\n*** You Won! ***\n", userChoice, computerOption)
		} else {
			fmt.Printf("You: %v\nComputer: %v.\n*** You Lose! ***\n", userChoice, computerOption)
		}
	}
}

func win(player string, opponent string) bool {
	// r > s, s > p, p > r
	result := false
	if (player == "r" && opponent == "s") || (player == "p" && opponent == "r") || (player == "s" && opponent == "p") {
		result = true
	}
	return result
}

func game() {
	user := ""
	computer := [3]string{"r", "p", "s"}
	rand.Seed(time.Now().UTC().UnixNano())
	computerRand := computer[rand.Intn(len(computer))]

	fmt.Println("What is your choice 'r' for rock, 'p', for paper and 's' for scissors")
	fmt.Scanln(&user)
	verification(user, computerRand)

}

func main() {
	game()
}
