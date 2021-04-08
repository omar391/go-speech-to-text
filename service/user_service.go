package service

import (
	"stt-service/models"
	"stt-service/repository"
	"stt-service/utils"
)

//login user
func Login(user_email string, user_password string) (result models.ApiBooleanResponse) {
	result.IsScuess = false
	if isLoginCredentialsValid(user_email, user_password) {
		result.IsScuess = true
		result.Msg = "Login successfull!"
	}
	result.Msg = "Incorrect login credentials! Please check your login info."
	return result
}

//Signup: add a new user to the repository
func AddNewUser(user *models.User) (result models.ApiBooleanResponse) {
	result.IsScuess = false

	if user.Email == "" && utils.IsEmailValid(user.Email) {
		result.Msg = "Please provide a valid email address!"
		return result

	} else if user.Password == "" {
		result.Msg = "Please provide valid password!"
		return result

	} else if user.Name == "" {
		result.Msg = "Please provide valid name!"
		return result

	} else if doesUserExists(user.Email) {
		result.Msg = "This email is already registered. Please provide new email!"
		return result
	}

	//convert password
	user.Password, _ = utils.GeneratePasswordHash(user.Password)

	repository.CreateUser(user)
	result.IsScuess = true
	return result
}

//check user info by email
func doesUserExists(user_email string) bool {
	user := repository.GetUserByEmail(user_email)
	return user.Email != ""
}

//check user info by email and password
func isLoginCredentialsValid(user_email string, password string) bool {

	user := repository.GetUserByEmail(user_email)
	return user.Email != "" && utils.ComparePasswordAndHash(password, user.Password)
}
