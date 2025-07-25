package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID        string         `json:"-" gorm:"primaryKey;default:gen_random_uuid()"`
	Title     string         `json:"title"`
	IsDone    bool           `json:"is_done"`
	UserID    uuid.UUID      `json:"user_id"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
