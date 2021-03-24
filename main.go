package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	url := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	now := time.Now().Format(time.RFC3339)
	fmt.Printf("Connect to %s as %s at %s\n", url, username, now)
	c, err := redis.Dial(
		"tcp", url+":6379",
		redis.DialUseTLS(true),
		redis.DialUsername(username),
		redis.DialPassword(password),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer c.Close()
	fmt.Printf("SET a: %s\n", now)
	if _, err = c.Do("SET", "a", now); err != nil {
		fmt.Println(err)
	}
	fmt.Print("GET a: ")
	v, err := c.Do("GET", "a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", string(v.([]byte)))
	}
	fmt.Print("INFO")
	v, err = c.Do("INFO")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", string(v.([]byte)))
	}
}
