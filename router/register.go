package router

import (
	"net/http"
	"stt-service/models"
	"stt-service/service"
	"stt-service/utils"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {

	email, fe := c.GetPostForm("email")
	pass, fp := c.GetPostForm("password")
	name, fn := c.GetPostForm("name")

	response := models.ApiResponse{}
	if !fe || !fp || !fn {
		response.IsSuccess = false
		response.Msg = "Please provide ALL registration credentials!"

	} else {
		var id uint
		response, id = service.AddNewUser(&models.User{Email: email, Password: pass, Name: name})
		if response.IsSuccess {
			response.Token, _ = utils.GenerateJWT(id)
			saveAuthSession(c, id, response.Token)
		}
	}
	c.JSON(http.StatusOK, response)
}
