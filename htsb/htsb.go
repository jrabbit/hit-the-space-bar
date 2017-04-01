package main

import (
	"fmt"
	// "bufio"
	// "os"
	"github.com/nsf/termbox-go"

)


func main() {
	fmt.Println("Welcome to hit the spacebar 2017 GOTY edition")
	// var x string
	// var score int
	// scanner := bufio.NewScanner(os.Stdin)

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
mainloop:
	for{
		return
	}

	// for scanner.Scan() {
		
	// 	x = scanner.Text()
	// 	if x == " " {
	// 		fmt.Println("space detected")
	// 		score += 1
	// 	} else {
	// 		fmt.Println()
	// 	}
	// 	fmt.Println("your score:", score)
	// }
}
