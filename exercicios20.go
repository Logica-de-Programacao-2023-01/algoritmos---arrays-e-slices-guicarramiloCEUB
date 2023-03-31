package main

import "fmt"

func main() {
	var slice = []int{}
	var valor, tamanho int
	fmt.Printf("tamanho da slice: ")
	fmt.Scan(&tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Printf("insira numero para array: ")
		fmt.Scan(&valor)
		slice = append(slice, valor)
	}
	crescente := 1
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			fmt.Println("Array nao é crescente")
			crescente = 0
			break
		}
	}
	if crescente == 1 {
		fmt.Println("array é crescente")
	}
}
