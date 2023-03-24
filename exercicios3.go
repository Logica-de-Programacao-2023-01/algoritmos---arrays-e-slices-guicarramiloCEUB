package main

import "fmt"

func main() {
	var array = [4]float64{1.1, 1.2, 1.3, 1.4}
	var produto = 1.0
	for i := 0; i < len(array); i++ {
		produto *= array[i]
	}
	fmt.Printf("%.2f", produto)
}
