package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Area struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Area) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Whs struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Prefix      string    `validate:"required,min=1,max=10" gorm:"not null;index;unique;size:10" json:"prefix" form:"prefix"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title" form:"title"`
	Value       string    `validate:"required,min=1,max=5" gorm:"size:5" json:"value" form:"value"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Whs) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Factory struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	InvPrefix   string    `validate:"required,min=1,max=10" gorm:"size:10" json:"inv_prefix" form:"inv_prefix"`
	LabelPrefix string    `validate:"required,min=1,max=10" gorm:"size:10" json:"label_prefix" form:"label_prefix"`
	Value       int64     `json:"value" form:"value"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Factory) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Unit struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Unit) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type ItemType struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *ItemType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Position struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Position) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Section struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Section) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Department struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Department) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Shipment struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Prefix      string    `validate:"required,min=1,max=10" gorm:"size:10" json:"prefix" form:"prefix"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Shipment) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type MailType struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	FactoryID   string    `json:"factory_id" form:"factory_id"`
	Prefix      string    `validate:"required,min=1,max=10" gorm:"not null;index;unique;size:50" json:"prefix" form:"prefix"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;size:50" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
	Factory     Factory   `gorm:"foreignKey:FactoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"factory"`
}

func (obj *MailType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type MailBox struct {
	ID        string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	AreaID    string    `json:"area_id" form:"area_id"`
	MailID    string    `gorm:"size:50;" json:"mail_id" form:"mail_id"`
	Password  string    `gorm:"size:50;" json:"password" form:"password"`
	Url       string    `json:"url" form:"url"`
	IsActive  bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" default:"now"`
	Area      Area      `gorm:"foreignKey:AreaID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"area"`
}

func (obj *MailBox) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type DownloadMailBox struct {
	ID           string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	MailBoxID    string    `json:"mail_box_id" form:"mail_box_id"`
	MailTypeID   string    `json:"mail_type_id" form:"mail_type_id"`
	BatchNo      string    `gorm:"not null;unique;size:50;" json:"batch_no" form:"batch_no"`
	Size         float64   `json:"size" form:"size" default:"0"`
	BatchID      string    `gorm:"not null;" json:"batch_id" form:"batch_id"`
	CreationDate time.Time `gorm:"type:date;" json:"creation_date" form:"creation_date" default:"now"`
	CreationTime time.Time `gorm:"type:time;" json:"creation_time" form:"creation_time" default:"now"`
	Flags        string    `gorm:"size:10;" json:"flags" form:"flags" default:"-"`
	Format       string    `gorm:"type:string;size:5;" json:"format" form:"format" default:"A"`
	Originator   string    `gorm:"type:string;size:25;" json:"originator" form:"originator" default:"-"`
	FilePath     string    `json:"file_path" form:"file_path"`
	IsDownload   bool      `gorm:"null" json:"is_download" form:"is_download" default:"false"`
	IsActive     bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt    time.Time `json:"created_at" default:"now"`
	UpdatedAt    time.Time `json:"updated_at" default:"now"`
	MailBox      MailBox   `gorm:"foreignKey:MailBoxID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"mail_box"`
	MailType     MailType  `gorm:"foreignKey:MailTypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"mail_type"`
}

func (obj *DownloadMailBox) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Part struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	Slug        string    `validate:"required,min=1,max=25" gorm:"size:25" json:"prefix" form:"prefix"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Part) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
