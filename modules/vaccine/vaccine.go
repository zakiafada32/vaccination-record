package vaccine

import (
	"time"

	"github.com/zakiafada32/vaccine/modules/user"
)

type Vaccine struct {
	ID             string    `gorm:"primaryKey;size:36"`
	Description    string    `gorm:"not null;size:255"`
	VaccinatedDate time.Time `gorm:"not null"`
	UserID         string    `gorm:"not null;size:36"`
	User           user.User
	HospitalID     string `gorm:"not null;size:36"`
	Hospital       Hospital
	DoctorID       string `gorm:"not null;size:36"`
	Doctor         Doctor
}

type Hospital struct {
	ID      string `gorm:"primaryKey;size:36"`
	Name    string `gorm:"not null;size:255"`
	Address string `gorm:"not null;size:255"`
}

type Doctor struct {
	ID        string `gorm:"primaryKey;size:36"`
	Name      string `gorm:"not null;size:255"`
	StrNumber string `gorm:"not null;size:20"`
}
