package model

import (
	"time"

	"github.com/google/uuid"
)

type Country struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
