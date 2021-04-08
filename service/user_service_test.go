package service

import (
	"fmt"
	"stt-service/models"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUserCreation(m *testing.T) {
	AddNewUser(&models.User{Email: "ding@dong.om", Name: "wow", Password: GeneratePasswordHash("")})
	s := []byte("")

	assert len(s)!=0
}
