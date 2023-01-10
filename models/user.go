package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	UserName  string    `validate:"required,min=5,max=10" gorm:"not null;column:username;index;unique;size:10" json:"username,omitempty" form:"username"`
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

func (obj *User) AfterCreate(tx *gorm.DB) (err error) {
	var log SystemLogger
	log.UserID = &obj.ID
	log.Title = "User created"
	log.Description = fmt.Sprintf("%s ลงทะเบียนเรียบร้อยแล้ว", obj.UserName)
	log.IsSuccess = true
	tx.Save(&log)
	return
}

type AuthSession struct {
	Header       string      `json:"header"`
	Prefix       interface{} `json:"prefix"`
	User         interface{} `json:"user_id"`
	Profile      interface{} `json:"profile"`
	Whs          interface{} `json:"whs"`
	Factory      interface{} `json:"factory"`
	Position     interface{} `json:"position"`
	Section      interface{} `json:"section"`
	Department   interface{} `json:"department"`
	AvatarURL    string      `json:"avatar_url"`
	EmployeeDate time.Time   `json:"employee_date"`
	JwtType      string      `json:"jwt_type"`
	JwtToken     string      `json:"jwt_token"`
	IsAdmin      bool        `json:"is_admin"`
}

type JwtToken struct {
	ID        string    `gorm:"primaryKey;size:60;" json:"id,omitempty"`
	UserID    *string   `gorm:"not null;unique;" json:"user_id,omitempty" form:"user_id" binding:"required"`
	Token     string    `gorm:"not null;unique;" json:"token,omitempty" form:"token"`
	IsActive  bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt time.Time `json:"updated_at,omitempty" default:"now"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}

func (obj *JwtToken) AfterSave(tx *gorm.DB) (err error) {
	var log SystemLogger
	log.UserID = obj.UserID
	log.Title = "User Login"
	log.Description = "เข้าสู่ระบบเรียบร้อยแล้ว"
	log.IsSuccess = true
	tx.Save(&log)
	return
}

type Administrator struct {
	ID        string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	UserID    *string   `gorm:"unique;" json:"user_id,omitempty" form:"user_id"`
	IsActive  bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt time.Time `json:"updated_at,omitempty" default:"now"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}

func (obj *Administrator) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
