package repository

import (
	"stt-service/models"
	"stt-service/utils"
)

// Migrate the schema on startup
func init() {
	db := utils.OpenSQLiteDB()
	db.AutoMigrate(models.User{})
}

// Create a new user
func CreateUser(user *models.User) {
	db := utils.OpenSQLiteDB()
	db.Create(user)

	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
}

//get user info
func GetUserByEmail(user_email string) models.User {
	db := utils.OpenSQLiteDB()

	user := &models.User{}
	db.First(user, "email = ?", user_email)
	return *user
}
