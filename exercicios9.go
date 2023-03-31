package main

import "fmt"

func main() {
	var slice = []float64{}
	var valor float64
	fmt.Printf("insira valor: ")
	fmt.Scan(&valor)
	for i := 0; i < 6; i++ {
		slice = append(slice, valor)
	}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%.2f ", slice[i])
	}
}
