package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type SystemLogger struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Title       string    `validate:"required,min=5,max=25" gorm:"size:25" json:"title,omitempty" form:"title"`
	Description string    `json:"description,omitempty" form:"description"`
	IsSuccess   bool      `gorm:"null" json:"is_success,omitempty" form:"is_success" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *SystemLogger) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
