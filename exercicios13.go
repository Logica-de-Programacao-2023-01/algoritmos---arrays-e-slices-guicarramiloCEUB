package main

import "fmt"

func main() {
	var array = [7]int{3, 5, 6, 7, 1, 2, 9}
	var valor int
	fmt.Printf("informe valor: ")
	fmt.Scan(&valor)
	array[0] += valor
	array[len(array)-1] += valor
	fmt.Printf("primeiro: %d\nultimo: %d", array[0], array[len(array)-1])

}
