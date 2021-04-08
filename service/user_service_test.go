package service

import (
	"stt-service/models"
	"testing"
)

func TestUserCreation(m *testing.T) {
	AddNewUser(&models.User{Email: "ding@dong.om", Name: "wow", Password: GeneratePasswordHash("")})
	// s := []byte("")

	// assert len(s)!=0
}
