package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        string    `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Desc      string    `json:"desc"`
	Price     int64     `gorm:"not null" json:"price"`
	UserID    string    `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.NewString()
	}
	return nil
}
