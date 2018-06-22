package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}

			if _, err = c.Do("SELECT", 0); err != nil {
				c.Close()
				return nil, err
			}

			return c, nil
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", "name", "majun", "EX", 10)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	name, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Println("redis get name:", name)
	}
}
