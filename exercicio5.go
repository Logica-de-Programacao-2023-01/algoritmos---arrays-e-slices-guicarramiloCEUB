package main

import "fmt"

func main() {
	var matriz [3][2]int
	var valor int
	for i := 0; i < 3; i++ {
		for y := 0; y < 2; y++ {
			fmt.Printf("insira valor para linha %d coluna %d: ", i+1, y+1)
			fmt.Scan(&valor)
			matriz[i][y] = valor
		}
	}
	for i := 0; i < 3; i++ {
		fmt.Println(matriz[i])
	}
}
