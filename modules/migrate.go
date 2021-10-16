package modules

import (
	"github.com/zakiafada32/vaccine/modules/user"
	"github.com/zakiafada32/vaccine/modules/vaccine"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&vaccine.Vaccine{},
		&vaccine.Hospital{},
		&vaccine.Doctor{},
	)
}
