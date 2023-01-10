package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Whs struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Prefix      string    `validate:"required,min=1,max=10" gorm:"not null;index;unique;size:10" json:"prefix,omitempty" form:"prefix"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title,omitempty" form:"title"`
	Value       string    `validate:"required,min=1,max=5" gorm:"size:5" json:"value,omitempty" form:"value"`
	Description string    `json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *Whs) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Unit struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title,omitempty" form:"title"`
	Description string    `json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *Unit) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Position struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title,omitempty" form:"title"`
	Description string    `json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *Position) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Section struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title,omitempty" form:"title"`
	Description string    `json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *Section) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Department struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title,omitempty" form:"title"`
	Description string    `json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *Department) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Shipment struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title,omitempty" form:"title"`
	Description string    `json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *Shipment) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
