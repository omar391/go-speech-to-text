package main

import (
	"log"
	"stt-service/conf"
	"stt-service/srv"
)

func main() {

	if err := srv.Start(conf.Config.Addr); err != nil {
		log.Println("start: ", err)
	}
}
