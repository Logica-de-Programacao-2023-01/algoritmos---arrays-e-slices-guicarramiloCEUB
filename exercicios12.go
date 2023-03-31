package main

import "fmt"

func main() {
	var array = [5]int{3, 6, 7, 2, 9}
	var slice = []int{}
	for i := 0; i < len(array); i++ {
		if array[i]%3 == 0 {
			slice = append(slice, array[i])
		}
	}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", slice[i])
	}
}
