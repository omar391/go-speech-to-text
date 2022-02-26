package router

import (
	"net/http"
	"stt-service/models"
	"stt-service/service"
	"stt-service/utils"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	email, fe := c.GetPostForm("email")
	pass, fp := c.GetPostForm("password")

	response := models.ApiResponse{}
	if !fe || !fp {
		response.IsSuccess = false
		response.Msg = "Please provide ALL login credentials!"

	} else {
		var id uint
		response, id = service.Login(email, pass)
		if response.IsSuccess {
			response.Token, _ = utils.GenerateJWT(id)
			saveAuthSession(c, id, response.Token)
		}
	}
	c.JSON(http.StatusOK, response)
}
