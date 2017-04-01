package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
)

// func main() {
// 	fmt.Println("Hello!")
// 	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
// 	})
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func root() string {
	return "<h1>Hit the spacebar scoreboard server</h1>"
}

func scoreboard() string {
	return "<body><ol><li>JEB - 9001</li></ol></body>"
}

func postscore(req *http.Request) string {
	fmt.Println(req)
	return "OK"
}

func main() {
	m := martini.Classic()
	m.Get("/", root)
	m.Get("/scoreboard", scoreboard)
	m.Post("/scoreboard/submit", postscore)
	m.Run()
}
