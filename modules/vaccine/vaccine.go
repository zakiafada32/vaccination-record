package vaccine

import (
	"time"

	"github.com/zakiafada32/vaccine/modules/user"
)

type Vaccine struct {
	ID             uint32    `gorm:"primaryKey"`
	Description    string    `gorm:"not null;size:255"`
	VaccinatedDate time.Time `gorm:"not null"`
	UserID         string    `gorm:"not null;size:36"`
	User           user.User
	HospitalID     string `gorm:"not null"`
	Hospital       Hospital
	DoctorID       string `gorm:"not null"`
	Doctor         Doctor
}

type Hospital struct {
	ID      uint32 `gorm:"primaryKey"`
	Name    string `gorm:"not null;size:255;unique"`
	Address string `gorm:"not null;size:255"`
}

type Doctor struct {
	ID        uint32 `gorm:"primaryKey"`
	Name      string `gorm:"not null;size:255"`
	StrNumber string `gorm:"not null;size:20;unique"`
}
