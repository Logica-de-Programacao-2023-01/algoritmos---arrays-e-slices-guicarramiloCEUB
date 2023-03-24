package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5}
	slice = append(slice[:2], slice[3:]...)
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
