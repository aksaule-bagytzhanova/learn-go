package main

import "fmt"

func main() {
	fmt.Println("counting")
	//last-in-first-out
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
