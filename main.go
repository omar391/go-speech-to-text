package main

import (
	"log"
	"stt-service/conf"
	"stt-service/router"
)

func main() {

	//starting the http routers with the address/port set in the config file
	if err := router.Start(conf.Config.ADDR); err != nil {
		log.Println("start: ", err)
	}
}
