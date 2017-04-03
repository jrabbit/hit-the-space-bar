package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/go-redis/redis"
	"net/http"
	"os"
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

func getBoard(client *redis.Client) map[string]int64 {
	scores := make(map[string]int64)
	iter := client.Scan(0, "", 0).Iterator()
	for iter.Next() {
		ret, err := client.Get(iter.Val()).Int64()
		if err != nil {
			panic(err)
		}
		scores[iter.Val()] = ret
		fmt.Println(iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return scores
}

func scoreboard(client *redis.Client) string {
	scores := getBoard(client)
	var lines string
	lines = "<html><body><ol><li> JEB -  9001 </li>"
	for key, value := range scores {
		lines += fmt.Sprintf("<li> %s -  %d </li>", key, value)
	}
	lines += "</ol></body></html>"
	return lines
}

func postScore(req *http.Request) string {
	form := req.PostForm
	fmt.Println(form)
	return "OK"
}

func main() {
	m := martini.Classic()
	redisURL := os.Getenv("REDIS_URL")
	var client *redis.Client
	if redisURL == "" {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	} else {
		opt, _ := redis.ParseURL(redisURL)
		client = redis.NewClient(opt)
	}

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// db := &client
	m.Map(client)
	m.Get("/", root)
	m.Get("/scoreboard", scoreboard)
	m.Post("/scoreboard/submit", postScore)
	m.Run()
}
