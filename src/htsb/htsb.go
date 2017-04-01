package main

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/nsf/termbox-go"
	"log"
)

func upload_score(score int) {
	// json := "{'score':5}"
	ro := &grequests.RequestOptions{JSON: map[string]int{"score": score}}

	resp, err := grequests.Post("http://httpbin.org/post", ro)
	// You can modify the request by passing an optional RequestOptions struct

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	fmt.Println(resp.String())
}

func cleanup(score *int, scoreboard bool) {
	fmt.Println("Thanks for playing!")
	fmt.Println("Final Score: ", *score) // dereference the score

	if scoreboard {
		upload_score(*score)
	}
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
