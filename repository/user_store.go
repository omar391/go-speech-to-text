package repository

import (
	"stt-service/models"
	"stt-service/utils"
)

// Migrate the schema on startup
func init() {
	db := utils.OpenSQLiteDB()
	db.AutoMigrate(&models.User{})
}

// Create a new user
func CreateUser(user *models.User) {
	db := utils.OpenSQLiteDB()
	db.Create(user)
}

//get user info
func GetUserByEmail(user_email string) models.User {
	db := utils.OpenSQLiteDB()

	user := &models.User{}
	db.First(user, "email = ?", user_email)
	return *user
}
