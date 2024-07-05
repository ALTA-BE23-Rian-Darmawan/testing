package encrypts_test

import (
	"BE23TODO/utils/encrypts"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	hashService := encrypts.NewHashService()

	t.Run("success", func(t *testing.T) {
		password := "password123"

		hashedPassword, err := hashService.HashPassword(password)

		assert.NoError(t, err)
		assert.NotEmpty(t, hashedPassword)
		assert.NotEqual(t, password, hashedPassword)
	})

	t.Run("failed", func(t *testing.T) {
		// Test case for error scenario when hashing an empty password
		emptyPassword := ""
		hashedPassword, err := hashService.HashPassword(emptyPassword)

		assert.NoError(t, err)
		assert.NotEmpty(t, hashedPassword)

		// Check if the hashed password is a valid bcrypt hash
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(emptyPassword))
		assert.NoError(t, err)
	})
}

func TestCheckPasswordHash(t *testing.T) {
	hashService := encrypts.NewHashService()

	t.Run("success", func(t *testing.T) {
		password := "password123"
		hashedPassword, err := hashService.HashPassword(password)
		assert.NoError(t, err)

		isValid := hashService.CheckPasswordHash(hashedPassword, password)

		assert.True(t, isValid)
	})

	t.Run("incorrect password", func(t *testing.T) {
		password := "password123"
		hashedPassword, err := hashService.HashPassword(password)
		assert.NoError(t, err)

		isValid := hashService.CheckPasswordHash(hashedPassword, "wrongpassword")

		assert.False(t, isValid)
	})

	t.Run("empty password", func(t *testing.T) {
		password := "password123"
		hashedPassword, err := hashService.HashPassword(password)
		assert.NoError(t, err)

		isValid := hashService.CheckPasswordHash(hashedPassword, "")

		assert.False(t, isValid)
	})
}
