package main

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"strconv"
	"strings"
)

func promptName() string {
	var name string
	fmt.Println("Whats your name?")
	fmt.Scanln(&name)
	name = strings.ToUpper(name)
	return name
}

func uploadScore(score int) {
	// TODO: Add name parsing/guessing/suggestions
	name := promptName()
	ro := &grequests.RequestOptions{Data: map[string]string{"score": strconv.Itoa(score), "name": name}}
	scoreHost, ok := os.LookupEnv("HTSB_SCOREBOARD")
	if !ok {
		scoreHost = "https://htsb.herokuapp.com/"
	}
	resp, err := grequests.Post(scoreHost+"/scoreboard/submit", ro)

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	fmt.Println(resp.String())
}

func cleanup(score *int, scoreboard bool) {
	fmt.Println("Thanks for playing!")
	fmt.Println("Final Score: ", *score) // dereference the score

	if scoreboard {
		uploadScore(*score)
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
