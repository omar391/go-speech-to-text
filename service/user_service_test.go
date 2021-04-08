package service

import (
	"stt-service/models"
	"testing"
)

//testing if we could create a new user
func TestAddNewUser(t *testing.T) {
	user := &models.User{Name: "Test 1", Email: "test@test.com", Password: "xyz123"}
	result := AddNewUser(user)
	if !result.IsScuess {
		t.Fatal("Creation failed!")
	}
}
