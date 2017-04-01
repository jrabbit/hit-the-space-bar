package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to hit the spacebar 2017 GOTY edition")
	var x string
	for {
		fmt.Scan(&x)
		if x == " " {
			fmt.Println("space detected")
		}
		fmt.Println(x)
	}
}
