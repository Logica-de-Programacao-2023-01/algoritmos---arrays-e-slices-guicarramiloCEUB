package main

import "fmt"

func main() {
	var slice = []int{}
	var valor, repetiu int
	for i := 0; i < 5; i++ {
		fmt.Printf("insira um numero para a slice: ")
		fmt.Scan(&valor)
		repetiu = 0
		for y := 0; y < len(slice); y++ {
			if repetiu == 0 {
				if valor == slice[y] {
					repetiu += 1
					fmt.Println("valor ja esta na lista")
					i--
					break
				}
			}
		}
		if repetiu == 0 {
			slice = append(slice, valor)
		}
	}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", slice[i])
	}
}
