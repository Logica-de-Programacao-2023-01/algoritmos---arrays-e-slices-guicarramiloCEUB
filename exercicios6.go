package main

import "fmt"

func main() {
	var lista = [10]int{10, 4, 7, 9, 23, 45, 21, 53, 22, 66}
	var numero int
	fmt.Printf("informe numero para busca: ")
	fmt.Scan(&numero)
	for i := 0; i < len(lista); i++ {
		if lista[i] == numero {
			fmt.Printf("numero é o %d da lista.", i+1)
			break
		}
	}
}
