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

	// str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTgyMTgyMTYsInVpZCI6MTR9.8oGBSN44SBfS__Rm2qtKtqqXk-zwTNWEsKEL4x81wpc"
	// b:=utils.VerifyJWT(str)

	// response := models.ApiBooleanResponse{}
	// email := "email@F." + strconv.FormatInt(time.Now().Unix(), 10)
	// response, id := service.AddNewUser(&models.User{Email: email, Password: "pass", Name: "name"})
	// if response.IsScuess {
	// 	response.Token, _ = utils.GenerateJWT(id)
	// 	//	response = response2
	// }
	// fmt.Println(response,b)

}
