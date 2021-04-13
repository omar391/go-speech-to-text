package main

import (
	"log"
	"os"
	"stt-service/conf"
	"stt-service/router"
)

func main() {
	//	starting the http routers with the address/port set in the config file
	addr := conf.Config.ADDR
	if os.Getenv("PORT") != "" {
		addr = ":" + addr
	}

	if err := router.Start(addr); err != nil {
		log.Println("start: ", err)
	}
}
