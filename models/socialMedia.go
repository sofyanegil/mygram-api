package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name           string    `gorm:"not null" json:"name"`
	SocialMediaUrl string    `gorm:"not null" json:"social_media_url"`
	UserID         string    `gorm:"index;references:ID"`
	CreatedAt      time.Time `gorm:"index"`
	UpdatedAt      time.Time `gorm:"index"`
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	sm.ID = uuid.New()
	return
}
