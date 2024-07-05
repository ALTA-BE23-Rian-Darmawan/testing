package datatodos

import "gorm.io/gorm"

type Todos struct {
	gorm.Model
	UserID      uint
	TodoName    string
	Description string
}
