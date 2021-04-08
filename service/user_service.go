package service

import (
	"crypto/rand"
	"log"
	"stt-service/models"
	"stt-service/repository"

	"golang.org/x/crypto/argon2"
)

// add a new user to the repository
func AddNewUser(user *models.User) {
	repository.CreateUser(user)
}

//get user info
func IsUserAvailable(user_email string) bool {
	user_info := repository.ReadUser(user_email)
	return user_info.Email != ""
}

//generate a secure hasg for a given password
func GeneratePasswordHash(password string) []byte {
	params := &models.PasswordParam{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	// Pass the plaintext password and parameters to our generateFromPassword
	// helper function.
	hash, err := generateFromPassword(password, params)
	if err != nil {
		log.Fatal(err)
	}

	return hash
}

// Generate a cryptographically secure random salt
func generateFromPassword(password string, p *models.PasswordParam) (hash []byte, err error) {
	salt, err := generateRandomBytes(p.SaltLength)
	if err != nil {
		return nil, err
	}

	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	hash = argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	return hash, nil
}

//get a random salt for a given length
func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
