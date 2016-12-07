package interfaces

import (
	"app"

	"github.com/jinzhu/gorm"
)

// InitDB creates tables
func InitDB(db *gorm.DB) error {
	return db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(
		&app.User{},
		&app.Account{},
		&app.Transaction{},
		&app.TransactionTag{},
	).Error
}
