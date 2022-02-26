package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"stt-service/conf"
	"stt-service/models"
	"stt-service/service"

	"github.com/gin-gonic/gin"
)

func transcribe(c *gin.Context) {
	response := models.ApiResponse{IsSuccess: false}
	token, fe := c.GetPostForm("token")
	is_save_file, _ := c.GetPostForm("is_save_file")
	re := regexp.MustCompile(conf.Config.ALLOWED_FILE_EXTS)

	var tmp_file *os.File
	var err error

	// get uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		response.Msg = err.Error()

	} else if !fe || !isSessionAuthenticated(c, token) {
		// check if valid token is available
		response.Msg = "No valid session found. Please login first."

	} else if file.Size > int64(conf.Config.MAX_FILE_UPLOAD_SIZE*1024*1024) {
		// check file size validity
		response.Msg = "File upload size exceeded! Max size is " + fmt.Sprint(conf.Config.MAX_FILE_UPLOAD_SIZE) + " MB"

	} else if re.FindStringIndex(file.Filename) == nil {
		// check file extension validity
		response.Msg = "File extension is not supported. Please from among these file type patterns: " + conf.Config.ALLOWED_FILE_EXTS

	} else if tmp_file, err = ioutil.TempFile("", "*"+file.Filename); err != nil {
		// creating temporary file
		response.Msg = "Error in creating temporary file in the disk!"

	} else if err = c.SaveUploadedFile(file, tmp_file.Name()); err != nil {
		// saving uploaded file to temporary file
		response.Msg = "Error in saving uploaded file in the disk!"

	} else {

		// get user id
		user_id_str, _ := getUserId(c)
		user_id, _ := strconv.ParseUint(user_id_str, 10, 32)

		// Transcribe the AUDIO file
		text, err := service.SaveTextFromAudio(file.Filename, tmp_file.Name(), uint(user_id), regexp.MustCompile(`(?i)\btrue\b`).FindStringIndex(is_save_file) != nil)

		// if no err this time then we are all good!
		if err == nil {
			response.IsSuccess = true
			response.Data = text
			response.Msg = "Speech to text was successfully saved!"
			response.Token = token

		} else {
			response.Msg = text + "; " + err.Error()
		}
	}

	c.JSON(http.StatusOK, response)
}
