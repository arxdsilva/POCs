package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/go-redis/redis"
)

func main() {
	url, err := url.Parse(os.Getenv("REDISCLOUD_URL"))
	fmt.Println("err: ", err)
	fmt.Printf("url: %+v\n", url.Host)
	fmt.Printf("url: %+v\n", url.User.Username())
	p, _ := url.User.Password()
	fmt.Printf("url: %+v\n", p)
	client := redis.NewClient(&redis.Options{
		Addr:     url.Host,
		Password: p,
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
