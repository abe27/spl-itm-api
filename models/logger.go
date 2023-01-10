package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type SystemLogger struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	UserID      *string   `json:"user_id" form:"user_id"`
	Title       string    `validate:"required,min=5,max=25" gorm:"size:25" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsSuccess   bool      `gorm:"null" json:"is_success" form:"is_success" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
	User        User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

func (obj *SystemLogger) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
