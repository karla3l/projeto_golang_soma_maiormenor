package main

import "fmt"

func main() {
	cont()
}
func cont() {
	var maior, menor int

	for i := 0; i < 5; i++ {
		num := 0
		fmt.Printf("Digite um numero \t")
		fmt.Scan(&num)

		if num > maior {
			maior = num
		}
		if num < menor {
			menor = num
		}
	}

	fmt.Printf("Maior numero \t")
	fmt.Println(maior)
	fmt.Printf("Menor numero \t")
	fmt.Println(menor)
}
