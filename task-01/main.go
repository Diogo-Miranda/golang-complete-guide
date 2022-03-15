package main

import "fmt"

func main() {
	arr := []int{2, 10, 5, 7, 52, 30, 25, 78, 100, 3}

	for _, number := range arr {
		if number%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
