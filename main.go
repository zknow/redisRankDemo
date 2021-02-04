package main

import (
	"log"

	"github.com/zknow/redisRankDemo/api/database"
	"github.com/zknow/redisRankDemo/api/router"
)

func main() {
	defer database.Conn.Close()
	if !database.Ping() {
		return
	}
	//serve on 8080 port
	r := router.GetRouter()
	if r.Run() != nil {
		log.Panicln("Router Listen Failed!")
	}
}
