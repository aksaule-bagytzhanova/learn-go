package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	for i := 7; i < 15; i++ {
		fmt.Println(i)
	}
}
