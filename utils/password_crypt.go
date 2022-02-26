package utils

import (
	"log"

	"github.com/alexedwards/argon2id"
)

// CreateHash returns a Argon2id hash of a plain-text password using the
// provided algorithm parameters. The returned hash follows the format used
// by the Argon2 reference C implementation and looks like this:
// $argon2id$v=19$m=65536,t=3,p=2$c29tZXNhbHQ$RdescudvJCsgt3ub+b+dWRWJTmaaJObG
func GeneratePasswordHash(password string) (hash string, err error) {
	hash, err = argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}

	return hash, err
}

// ComparePasswordAndHash performs a constant-time comparison between a
// plain-text password and Argon2id hash, using the parameters and salt
// contained in the hash. It returns true if they match, otherwise it returns
// false.
func ComparePasswordAndHash(password, encodedHash string) (match bool) {
	match, err := argon2id.ComparePasswordAndHash(password, encodedHash)
	if err != nil {
		log.Fatal(err)
	}

	return match
}
