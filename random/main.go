package main

import (
	"fmt"
	"math/rand"
)

func guess(x int) {
	max := 10
	min := 0
	//essa variavel é criada e o primeiro valor alocado esta se mantendo. Resolver isso (provavel que esta assim devido a usar a marmotinha para atribuir valor)
	random_number := (rand.Intn(max-min) + min)

	guess := 0
	for guess != random_number {
		fmt.Println("Guess the number: ")
		fmt.Scanln(&guess)

		if guess < random_number {
			fmt.Println("Sorry, guess again. Too low")
		} else if guess > random_number {
			fmt.Println("Sorry, guess again. Too High")
		}
	}
	fmt.Printf("Yes, are right! The number is %d", random_number)
}

func main() {
	guess(10)
}
