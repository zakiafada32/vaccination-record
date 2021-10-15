package modules

import (
	"github.com/zakiafada32/vaccine/modules/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
	)
}
