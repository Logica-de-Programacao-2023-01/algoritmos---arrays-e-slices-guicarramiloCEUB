package main

import "fmt"

func main() {
	var array = [5]int{1, 3, 3, 4, 5}
	var slice = []int{}
	var valor int
	fmt.Printf("insira valor: ")
	fmt.Scan(&valor)
	for i := 0; i < len(array); i++ {
		if array[i] != valor {
			slice = append(slice, array[i])
		}
	}

	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", slice[i])
	}
}
