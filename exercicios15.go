package main

import "fmt"

func main() {
	var array = [10]float64{1.4, 2.6, 4.7, 5.4, 3.7, 6.7, 3.5, 6.9, 7.7, 11.2}
	var slice = []float64{}
	for i := 0; i < len(array); i++ {
		if array[i] >= 5 {
			slice = append(slice, array[i])
		}
	}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%.2f ", slice[i])
	}
}
