package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID           string     `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	UserName     string     `validate:"required,min=5,max=10" gorm:"not null;column:username;index;unique;size:10" json:"username" form:"username"`
	FirstName    string     `gorm:"size:100" json:"firstname" form:"firstname"`
	LastName     string     `gorm:"size:100" json:"lastname" form:"lastname"`
	Email        string     `validate:"required,email,min=15,max=50" gorm:"not null;unique;size:50;" json:"email" form:"email"`
	Password     string     `validate:"required,min=6,max=60" gorm:"not null;size:60;" json:"-" form:"password"`
	AreaID       *string    `json:"area_id" form:"area_id"`
	WhsID        *string    `json:"whs_id" form:"whs_id"`
	FactoryID    *string    `json:"factory_id" form:"factory_id"`
	PositionID   *string    `json:"position_id" form:"position_id"`
	SectionID    *string    `json:"section_id" form:"section_id"`
	DepartmentID *string    `json:"department_id" form:"department_id"`
	AvatarURL    string     `json:"avatar_url" form:"avatar_url"`
	IsActive     bool       `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt    time.Time  `json:"created_at" default:"now"`
	UpdatedAt    time.Time  `json:"updated_at" default:"now"`
	Area         Area       `gorm:"foreignKey:AreaID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"area"`
	Whs          Whs        `gorm:"foreignKey:WhsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"whs"`
	Factory      Factory    `gorm:"foreignKey:FactoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"factory"`
	Position     Position   `gorm:"foreignKey:PositionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"position"`
	Section      Section    `gorm:"foreignKey:SectionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"section"`
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"department"`
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
	ID        string    `gorm:"primaryKey;size:60;" json:"id"`
	UserID    *string   `gorm:"not null;unique;" json:"user_id" form:"user_id" binding:"required"`
	Token     string    `gorm:"not null;unique;" json:"token" form:"token"`
	IsActive  bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" default:"now"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
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
	ID        string    `gorm:"primaryKey;size:21;" json:"id"`
	UserID    *string   `gorm:"unique;" json:"user_id" form:"user_id"`
	IsActive  bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" default:"now"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

func (obj *Administrator) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
