package migrations

import (
	datatodos "BE23TODO/features/Todos/dataTodos"
	datausers "BE23TODO/features/Users/dataUsers"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&datausers.Users{})
	db.AutoMigrate(&datatodos.Todos{})
}
