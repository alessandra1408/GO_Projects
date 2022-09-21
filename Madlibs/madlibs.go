package main

import "fmt"

func main() {
	var youtube string
	fmt.Println("Insira o nome do seu canal: ")
	fmt.Scanln(&youtube)
	fmt.Println("Subscribe to ", youtube)

}
