package main

import "fmt"

func main() {
	fmt.Println("Insira o nome do seu canal: ")
	var youtube string

	fmt.Scanln(&youtube)
	fmt.Println("\nSubscribe to ", youtube)
	fmt.Printf("\nSubscribe to %s ", youtube)
}
