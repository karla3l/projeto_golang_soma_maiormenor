package main

import "fmt"

func main() {
	soma()
}
func soma() {
	soma := 0
	for i := 0; i < 20; i++ {
		num := 0
		fmt.Printf("Digite um numero")
		fmt.Scan(&num)
		soma = num + soma
	}

	fmt.Println(soma)

}
