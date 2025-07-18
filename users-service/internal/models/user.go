package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	//Tasks     []Task    `json:"tasks" gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
