package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	// "net/http"
)

func upload_score(score int) {
	return
}

func cleanup(score *int, scoreboard bool) {
	fmt.Println("Thanks for playing!")
	fmt.Println("Final Score: ", *score) // dereference the score
	// fmt.Println("Wanna upload to the scoreboard?")
	// var x string
	// fmt.Scan(&x)
	if scoreboard {
		upload_score(*score)
	}
	// switch x {
	// case "Y":
	// 	fmt.Println("Ok uploading score")
	// 	return
	// case "N":
	// 	return
	// }
}

func main() {
	var score int
	scoreboard := true // implying
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	fmt.Println("Welcome to hit the spacebar 2017 GOTY edition")

	defer cleanup(&score, scoreboard)
	defer termbox.Close()
	// termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeySpace {
				score += 1
				fmt.Println("Your score: ", score)
			} else {
				break mainloop
			}
		}
	}
}
