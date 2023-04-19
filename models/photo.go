package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title     string    `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption   string    `json:"caption" form:"caption"`
	PhotoUrl  string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo_url is required"`
	UserID    string    `gorm:"index;references:ID"`
	Comments  []Comment `gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
