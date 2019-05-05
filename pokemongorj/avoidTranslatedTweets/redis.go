package main

import (
	"errors"
	"log"

	"github.com/go-redis/redis"
)

var (
	ErrKeyNotFound = errors.New("Redis does not contain key")
)

type Client struct {
	cli *redis.Client
}

type Config struct {
	Addr       string
	Password   string
	DB         int
	MaxRetries int
}

func NewClient(config *Config) (c *Client, err error) {
	cli := redis.NewClient(&redis.Options{
		Addr:       config.Addr,
		Password:   config.Password,
		DB:         0, //DEFAULT
		MaxRetries: config.MaxRetries,
	})
	c = &Client{cli}
	pong, err := cli.Ping().Result()
	log.Println("pong:", pong)
	log.Println("error:", err)
	return
}

func (client *Client) Find(id int64) (r string, err error) {
	cli := client.cli
	r, err = cli.Get(string(id)).Result()
	if err == redis.Nil {
		return r, ErrKeyNotFound
	}
	return
}

func (client *Client) Set(tweetID int64) (err error) {
	cli := client.cli
	return cli.Set(string(tweetID), ".", 0).Err()
}
