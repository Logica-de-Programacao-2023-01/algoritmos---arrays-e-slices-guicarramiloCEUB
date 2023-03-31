package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var i, y int
	var primeiro int
	fmt.Printf("insira indice do primeiro numero: ")
	fmt.Scan(&i)
	fmt.Printf("insira indice do segundo numero: ")
	fmt.Scan(&y)
	primeiro = slice[i]
	slice[i] = slice[y]
	slice[y] = primeiro
	fmt.Printf("valor do indice %d: %d\nvalor do indice %d: %d", i, slice[i], y, slice[y])
}
