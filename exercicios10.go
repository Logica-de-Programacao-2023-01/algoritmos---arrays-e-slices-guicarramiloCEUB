package main

import "fmt"

func main() {
	var slice = []int{2, 4, 5, 6, 17, 33, 56, 3, 5, 8}
	menor := slice[0]
	maior := slice[0]
	for i := 0; i < len(slice); i++ {
		if slice[i] >= maior {
			maior = slice[i]
		} else if slice[i] <= menor {
			menor = slice[i]
		}
	}
	fmt.Printf("maior: %d\nmenor: %d", maior, menor)
}
