package main

import "fmt"

func main() {
	var array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice = []int{}
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			slice = append(slice, array[i])
		}
	}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", slice[i])
	}
}
