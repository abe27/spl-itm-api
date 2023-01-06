package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	UserName  string    `validate:"required,min=5,max=10" gorm:"not null;column:username;index;unique;size:10" json:"user_name,omitempty" form:"user_name"`
	Email     string    `validate:"required,email,min=15,max=50" gorm:"not null;unique;size:50;" json:"email,omitempty" form:"email"`
	Password  string    `validate:"required,min=6,max=60" gorm:"not null;size:60;" json:"-" form:"password"`
	IsActive  bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
