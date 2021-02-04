package database

import (
	"log"

	"github.com/go-redis/redis"
)

var Conn *redis.Client

//DB redis instance
func init() {
	Conn = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Ping() bool {
	//try ping
	_, err := Conn.Ping().Result()
	if err != nil {
		log.Println("[Error] Redis connect failed ! :", err)
		return false
	}
	return true
}
