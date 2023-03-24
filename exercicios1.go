package main

import "fmt"

func main() {
	var array = []int{1, 2, 3}
	var sum = 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	fmt.Println(sum)
}
