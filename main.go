package main

import (
	"log"
	"stt-service/models"
	"stt-service/service"
)

func main() {

	// if err := srv.Start(conf.Config.Addr); err != nil {
	// 	log.Println("start: ", err)
	// }

	user := &models.User{Name: "Test 1", Email: "test@test.com", Password: "xyz123"}
	result := service.Login(user.Email, user.Password)
	if !result.IsScuess {
		log.Fatal("Creation failed!")
	} else {
		log.Println("logged in successfully!")
	}
}
