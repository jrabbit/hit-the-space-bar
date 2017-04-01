package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func main() {
	fmt.Println("Welcome to hit the spacebar 2017 GOTY edition")
	var score int

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeySpace {
				fmt.Println("Got spacebar!")
				score += 1
				fmt.Println("Your score: ", score)
			} else {
				break mainloop
			}
		}
	}
}
