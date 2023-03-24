package main

import "fmt"

func main() {
	var slice []int
	var soma, tamanho, num int
	fmt.Println("informe tamanho da lista: ")
	fmt.Scan(&tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Scan(&num)
		slice = append(slice, num)
		soma += num
	}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	fmt.Println(soma)
}
