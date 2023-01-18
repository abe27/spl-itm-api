package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Area struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
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
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
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
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
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
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
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
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
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
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
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
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
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
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
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
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
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

type ReceiveType struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	WhsID       string    `json:"whs_id" form:"whs_id"`
	Prefix      string    `validate:"required,min=1,max=10" gorm:"not null;index;unique;size:50" json:"prefix" form:"prefix"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;size:50" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
	Whs         Whs       `gorm:"foreignKey:WhsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"whs"`
}

func (obj *ReceiveType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type MailType struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
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
	ID        string    `gorm:"primaryKey;size:21;" json:"id"`
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
	ID           string    `gorm:"primaryKey;size:21;" json:"id"`
	MailBoxID    string    `json:"mail_box_id" form:"mail_box_id"`
	MailTypeID   string    `json:"mail_type_id" form:"mail_type_id"`
	BatchNo      string    `gorm:"not null;unique;size:50;" json:"batch_no" form:"batch_no"`
	Size         float64   `json:"size" form:"size" default:"0"`
	BatchID      string    `gorm:"not null;" json:"batch_id" form:"batch_id"`
	CreationDate time.Time `json:"creation_date" form:"creation_date" default:"now"`
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
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Slug        string    `validate:"required,min=1,max=25" gorm:"size:25;unique;" json:"slug" form:"slug"`
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

type Ledger struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	WhsID       *string   `gorm:"not null;" json:"whs_id" form:"whs_id" binding:"required"`
	AreaID      *string   `gorm:"not null;" json:"area_id" form:"area_id" binding:"required"`
	FactoryID   *string   `gorm:"not null;" json:"factory_id" form:"factory_id" binding:"required"`
	PartID      *string   `gorm:"not null;" json:"part_id" form:"part_id" binding:"required"`
	ItemTypeID  *string   `gorm:"not null;" json:"item_type_id" form:"item_type_id" binding:"required"`
	UnitID      *string   `gorm:"not null;" json:"unit_id" form:"unit_id" binding:"required"`
	DimWidth    float64   `json:"dim_width" form:"dim_width" default:"0"`
	DimLength   float64   `json:"dim_length" form:"dim_length" default:"0"`
	DimHeight   float64   `json:"dim_height" form:"dim_height" default:"0"`
	GrossWeight float64   `json:"gross_weight" form:"gross_weight" default:"0"`
	NetWeight   float64   `json:"net_weight" form:"net_weight" default:"0"`
	Qty         float64   `json:"qty" form:"qty" default:"0"`
	Ctn         float64   `json:"ctn" form:"ctn" default:"0"`
	IsActive    bool      `json:"is_active" form:"is_active" default:"true"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
	Whs         Whs       `gorm:"foreignKey:WhsID;references:ID" json:"whs"`
	Area        Area      `gorm:"foreignKey:AreaID;references:ID" json:"area"`
	Factory     Factory   `gorm:"foreignKey:FactoryID;references:ID" json:"factory"`
	Part        Part      `gorm:"foreignKey:PartID;references:ID;" json:"part"`
	ItemType    ItemType  `gorm:"foreignKey:ItemTypeID;references:ID" json:"part_type"`
	Unit        Unit      `gorm:"foreignKey:UnitID;references:ID" json:"unit"`
}

func (u *Ledger) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}

type Affcode struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Affcode) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Customer struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type ReviseOrder struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:5" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *ReviseOrder) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Pc struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:5" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Pc) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Commercial struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:5" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Commercial) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type SampleFlg struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:5" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *SampleFlg) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type OrderType struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:5" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *OrderType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type OrderZone struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Value       int64     `gorm:"not null;" json:"value" form:"value" binding:"required"`
	FactoryID   *string   `gorm:"not null;" json:"factory_id" form:"factory_id" binding:"required"`
	WhsID       *string   `gorm:"not null;" json:"whs_id" form:"whs_id" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
	Factory     Factory   `gorm:"foreignKey:FactoryID;references:ID" json:"factory"`
	Whs         Whs       `gorm:"foreignKey:WhsID;references:ID" json:"whs"`
}

func (obj *OrderZone) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type LastInvoice struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	FactoryID   *string   `gorm:"not null;" json:"factory_id" form:"factory_id" binding:"required"`
	AffcodeID   *string   `gorm:"not null;" json:"affcode_id" form:"affcode_id" binding:"required"`
	OnYear      int64     `gorm:"not null;" json:"on_year" form:"on_year" binding:"required"`
	LastRunning int64     `json:"last_running" form:"last_running" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
	Factory     Factory   `gorm:"foreignKey:FactoryID;references:ID" json:"factory"`
	Affcode     Affcode   `gorm:"foreignKey:AffcodeID;references:ID" json:"affcode"`
}

func (obj *LastInvoice) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type OrderGroupType struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:5" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *OrderGroupType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type OrderGroup struct {
	ID               string          `gorm:"primaryKey;size:21" json:"id"`
	UserID           *string         `gorm:"not null;" json:"user_id" form:"user_id" binding:"required"`
	AffcodeID        *string         `gorm:"not null;" json:"affcode_id" form:"affcode_id" binding:"required"`
	CustomerID       *string         `gorm:"not null;" json:"customer_id" form:"customer_id" binding:"required"`
	OrderGroupTypeID *string         `gorm:"not null;" json:"order_group_type_id" form:"order_group_type_id" binding:"required"`
	SubOrder         string          `gorm:"not null;size:15" json:"sub_order" form:"sub_order" binding:"required"`
	Description      string          `json:"description" form:"description" binding:"required"`
	IsActive         bool            `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt        time.Time       `json:"created_at" default:"now"`
	UpdatedAt        time.Time       `json:"updated_at" default:"now"`
	User             *User           `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Affcode          *Affcode        `gorm:"foreignKey:AffcodeID;references:ID" json:"affcode"`
	Customer         *Customer       `gorm:"foreignKey:CustomerID;references:ID" json:"customer"`
	OrderGroupType   *OrderGroupType `gorm:"foreignKey:OrderGroupTypeID;references:ID" json:"order_group_type"`
}

func (obj *OrderGroup) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Shelve struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `validate:"required,min=5,max=25" gorm:"not null;index;unique;size:25" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Shelve) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
