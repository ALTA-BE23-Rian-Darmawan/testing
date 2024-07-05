package datausers

import (
	datatodos "BE23TODO/features/Todos/dataTodos"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	FullName    string
	Email       string `gorm:"unique"`
	Password    string
	PhoneNumber string
	Address     string
	Todos       []datatodos.Todos `gorm:"foreignKey:UserID"`
}
