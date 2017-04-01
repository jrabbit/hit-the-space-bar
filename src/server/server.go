package main

import (
	"github.com/go-martini/martini"
)

// func main() {
// 	fmt.Println("Hello!")
// 	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
// 	})
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func root() string {
	return "Hello World"
}

func main() {
	m := martini.Classic()
	m.Get("/", root)
	m.Run()
}
