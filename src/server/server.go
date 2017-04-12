package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/macaron.v1"
	"html"
	"net/http"
	"os"
	"strconv"
)

func index(ctx *macaron.Context) {
	ctx.HTML(200, "home")
}

func zScore(client *redis.Client) string {
	vals, err := client.ZRevRangeByScoreWithScores("scoreboard", redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()
	if err != nil {
		panic(err)
	}
	var lines string
	lines = "<html><body><ol><li> JEB -  9001 </li>"
	for place, z := range vals {
		safeHTML := html.EscapeString(z.Member.(string))
		lines += fmt.Sprintf("<li> %s -  %d </li>", safeHTML, int(z.Score))
		fmt.Println(place)
	}
	lines += "</ol></body></html>"
	return lines
}

func postScore(client *redis.Client, req *http.Request) string {
	// form := req.PostForm
	req.ParseForm()
	fmt.Println(req.PostForm)
	score, err := strconv.Atoi(req.FormValue("score")) // this is dangerous
	if err != nil {
		// so they probably didn't send a score here
		panic(err)
	}
	resp := client.ZAdd("scoreboard", redis.Z{float64(score), req.FormValue("name")})
	fmt.Println(resp)
	return "OK"
}

func main() {
	m := macaron.Classic()
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
	m.Use(macaron.Renderer())
	m.Map(client)
	m.Get("/", index)
	m.Get("/scoreboard", zScore)
	m.Post("/scoreboard/submit", postScore)
	m.Run()
}
