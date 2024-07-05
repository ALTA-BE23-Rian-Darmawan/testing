package service_test

import (
	users "BE23TODO/features/Users"
	"BE23TODO/features/Users/service"
	"BE23TODO/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	mockUserData := new(mocks.DataUserInterface)
	mockHashService := new(mocks.HashInterface)
	userService := service.New(mockUserData, mockHashService)

	t.Run("success", func(t *testing.T) {
		user := users.User{
			FullName:    "John Doe",
			Email:       "johndoe@example.com",
			Password:    "password123",
			PhoneNumber: "08123456789",
			Address:     "Some Address",
		}

		mockHashService.On("HashPassword", mock.Anything).Return("hashedpassword", nil).Once()
		mockUserData.On("CreateAccount", mock.Anything).Return(nil).Once()

		err := userService.RegistrasiAccount(user)

		assert.NoError(t, err)
		mockHashService.AssertExpectations(t)
		mockUserData.AssertExpectations(t)
	})

	t.Run("validation error", func(t *testing.T) {
		user := users.User{}

		err := userService.RegistrasiAccount(user)

		assert.Error(t, err)
		assert.Equal(t, "[validation] nama/email/password/phone/address tidak boleh kosong", err.Error())
	})

	t.Run("hash error", func(t *testing.T) {
		user := users.User{
			FullName:    "John Doe",
			Email:       "johndoe@example.com",
			Password:    "password123",
			PhoneNumber: "08123456789",
			Address:     "Some Address",
		}

		mockHashService.On("HashPassword", mock.Anything).Return("", errors.New("hash error")).Once()

		err := userService.RegistrasiAccount(user)

		assert.Error(t, err)
		assert.Equal(t, "hash error", err.Error())
		mockHashService.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	mockUserData := new(mocks.DataUserInterface)
	mockHashService := new(mocks.HashInterface)
	userService := service.New(mockUserData, mockHashService)

	t.Run("success", func(t *testing.T) {
		email := "johndoe@example.com"
		password := "password123"
		user := &users.User{
			UserID:   1,
			Email:    email,
			Password: "hashedpassword",
		}

		mockUserData.On("AccountByEmail", email).Return(user, nil).Once()
		mockHashService.On("CheckPasswordHash", "hashedpassword", password).Return(true).Once()

		actualToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjAxNzAxNzksInVzZXJJZCI6MX0.bLgbBHU9P3BmP4rfEwfTIeINpOc_zFIDS6Ylr2odPcU"

		mockHashService.On("CreateToken", user.UserID).Return(actualToken, nil).Once()

		returnedUser, returnedToken, err := userService.LoginAccount(email, password)

		assert.NoError(t, err)
		assert.Equal(t, user, returnedUser)
		assert.Equal(t, actualToken, returnedToken)
		mockUserData.AssertExpectations(t)
		mockHashService.AssertExpectations(t)
	})

	t.Run("account not found", func(t *testing.T) {
		email := "johndoe@example.com"
		password := "password123"

		mockUserData.On("AccountByEmail", email).Return(nil, errors.New("account not found")).Once()

		_, _, err := userService.LoginAccount(email, password)

		assert.Error(t, err)
		assert.Equal(t, "account not found", err.Error())
		mockUserData.AssertExpectations(t)
	})

	t.Run("invalid password", func(t *testing.T) {
		email := "johndoe@example.com"
		password := "wrongpassword"
		user := &users.User{
			UserID:   1,
			Email:    email,
			Password: "hashedpassword",
		}

		mockUserData.On("AccountByEmail", email).Return(user, nil).Once()
		mockHashService.On("CheckPasswordHash", "hashedpassword", password).Return(false).Once()

		_, _, err := userService.LoginAccount(email, password)

		assert.Error(t, err)
		assert.Equal(t, "[validation] email atau password tidak sesuai", err.Error())
		mockUserData.AssertExpectations(t)
		mockHashService.AssertExpectations(t)
	})

	t.Run("token creation error", func(t *testing.T) {
		email := "johndoe@example.com"
		password := "password123"
		user := &users.User{
			UserID:   1,
			Email:    email,
			Password: "hashedpassword",
		}

		mockUserData.On("AccountByEmail", email).Return(user, nil).Once()
		mockHashService.On("CheckPasswordHash", "hashedpassword", password).Return(true).Once()

		// Mock CreateToken to simulate an error
		mockHashService.On("CreateToken", user.UserID).Return("", errors.New("token creation error")).Once()

		_, _, err := userService.LoginAccount(email, password)

		assert.Error(t, err)
		assert.Equal(t, "token creation error", err.Error())
		mockUserData.AssertExpectations(t)
		mockHashService.AssertExpectations(t)
	})
}
