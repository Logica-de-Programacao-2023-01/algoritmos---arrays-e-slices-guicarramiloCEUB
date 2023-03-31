package main

import "fmt"

func main() {
	var matriz = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var i, y int
	fmt.Printf("insira valor da linha: ")
	fmt.Scan(&i)
	fmt.Printf("insira valor da coluna: ")
	fmt.Scan(&y)
	fmt.Printf("valor armazenado no indice %d, %d: %d", i, y, matriz[i-1][y-1])
}
