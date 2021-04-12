package router

import (
	"net/http"
	"strconv"
	"stt-service/models"
	"stt-service/service"

	"github.com/gin-gonic/gin"
)

// Get all the uploaded transcribed data from current logged-in user
func search_all(c *gin.Context) {
	page, fp := c.GetPostForm("page_no")
	token, fe := c.GetPostForm("token")
	var page_no int
	var err error

	response := models.ApiResponse{}
	if !fe || !fp {
		response.IsSuccess = false
		response.Msg = "Please provide ALL registration credentials!"

	} else if !fe || !isSessionAuthenticated(c, token) {
		// check if valid token is available
		response.Msg = "No valid session token found. Please login first."

	} else if page_no, err = strconv.Atoi(page); err != nil {
		// Invalid page no value
		response.Msg = "Invalid page no value! Please check the input!"

	} else {

		// get user id
		user_id_str, _ := getUserId(c)
		user_id, _ := strconv.ParseUint(user_id_str, 10, 32)

		// Get audio text list
		response.Data = service.GetAudioDataByPage(page_no, 100, uint(user_id))
		response.IsSuccess = true
		response.Token = token
	}

	c.JSON(http.StatusOK, response)
}

// Get filtered data from search term from current logged-in user
func filter(c *gin.Context) {
	page, fp := c.GetPostForm("page_no")
	search_term, ft := c.GetPostForm("query")
	token, fe := c.GetPostForm("token")
	var page_no int
	var err error

	response := models.ApiResponse{}
	if !fe || !fp || !ft {
		response.IsSuccess = false
		response.Msg = "Please provide ALL registration credentials!"

	} else if !fe || !isSessionAuthenticated(c, token) {
		// check if valid token is available
		response.Msg = "No valid session token found. Please login first."

	} else if page_no, err = strconv.Atoi(page); err != nil {
		// Invalid page no value
		response.Msg = "Invalid page no value! Please check the input!"

	} else {

		// get user id
		user_id_str, _ := getUserId(c)
		user_id, _ := strconv.ParseUint(user_id_str, 10, 32)

		// Get audio text list
		response.Data = service.FilterAudioData(page_no, 100, search_term, uint(user_id))
		response.IsSuccess = true
		response.Token = token
	}

	c.JSON(http.StatusOK, response)
}
