package datausers

import (
	users "BE23TODO/features/Users"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.DataUserInterface {
	return &userQuery{
		db: db,
	}
}

// CreateAccount implements users.DataUserInterface.
func (u *userQuery) CreateAccount(account users.User) error {
	userGorm := Users{
		FullName:    account.FullName,
		Email:       account.Email,
		Password:    account.Password,
		PhoneNumber: account.PhoneNumber,
		Address:     account.Address,
	}
	tx := u.db.Create(&userGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// AccountByEmail implements users.DataUserInterface.
func (u *userQuery) AccountByEmail(email string) (*users.User, error) {
	var userData Users
	tx := u.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping
	var users = users.User{
		UserID:      userData.ID,
		FullName:    userData.FullName,
		Email:       userData.Email,
		Password:    userData.Password,
		Address:     userData.Address,
		PhoneNumber: userData.PhoneNumber,
	}

	return &users, nil
}
