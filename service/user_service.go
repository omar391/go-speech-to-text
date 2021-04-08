package service

import (
	"stt-service/models"
	"stt-service/repository"
	"stt-service/utils"
)

//login user
func Login(user_email string, user_password string) (result models.ApiBooleanResponse, user_id uint) {
	result.IsScuess = false
	user_id = 0
	if found, id := checkLoginCredentials(user_email, user_password); found {
		result.IsScuess = true
		result.Msg = "Login successfull!"
		user_id = id
	}
	result.Msg = "Incorrect login credentials! Please check your login info."
	return result, user_id
}

//Signup: add a new user to the repository
func AddNewUser(user *models.User) (result models.ApiBooleanResponse, id uint) {
	result.IsScuess = false
	id = 0

	if user.Email == "" && utils.IsEmailValid(user.Email) {
		result.Msg = "Please provide a valid email address!"
		return result, 0

	} else if user.Password == "" {
		result.Msg = "Please provide valid password!"
		return result, 0

	} else if user.Name == "" {
		result.Msg = "Please provide valid name!"
		return result, 0

	} else if doesUserExists(user.Email) {
		result.Msg = "This email is already registered. Please provide new email!"
		return result, 0
	}

	//convert password
	user.Password, _ = utils.GeneratePasswordHash(user.Password)

	repository.CreateUser(user)
	result.IsScuess = true
	return result, user.ID
}

//check user info by email
func doesUserExists(user_email string) bool {
	user := repository.GetUserByEmail(user_email)
	return user.Email != ""
}

//check user info by email and password
func checkLoginCredentials(user_email string, password string) (bool, uint) {

	user := repository.GetUserByEmail(user_email)
	return user.Email != "" && utils.ComparePasswordAndHash(password, user.Password), user.ID
}
