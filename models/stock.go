package models

import (
	"net"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type FrmCarton struct {
	WhsID    string  `json:"whs_id" form:"whs_id" binding:"required"`
	LedgerID string  `json:"ledger_id" form:"ledger_id" binding:"required"`
	SerialNo string  `json:"serial_no" form:"serial_no"`
	LotNo    string  `json:"lot_no" form:"lot_no"`
	LineNo   string  `json:"line_no" form:"line_no"`
	ReviseNo string  `json:"revise_no" form:"revise_no"`
	PalletNo string  `json:"pallet_no" form:"pallet_no"`
	ShelveID *string `json:"shelve_id" form:"shelve_id" binding:"required"`
	Qty      float64 `json:"qty" form:"qty" default:"0"`
	IsActive bool    `json:"is_active" form:"is_active" default:"true"`
}

type Carton struct {
	ID              string        `gorm:"primaryKey;size:21" json:"id"`
	LedgerID        *string       `gorm:"not null;" json:"ledger_id" form:"ledger_id" binding:"required"`
	ReceiveDetailID *string       `json:"receive_detail_id" form:"receive_detail_id" binding:"required"`
	SerialNo        string        `gorm:"not null;unique;size:15;" json:"serial_no" form:"serial_no"`
	LotNo           string        `gorm:"not null;size:15;" json:"lot_no" form:"lot_no"`
	LineNo          string        `gorm:"size:15;" json:"line_no" form:"line_no"`
	ReviseNo        string        `gorm:"size:15;" json:"revise_no" form:"revise_no"`
	PalletNo        string        `gorm:"size:15;" json:"pallet_no" form:"pallet_no"`
	ShelveID        *string       `gorm:"not null;" json:"shelve_id" form:"shelve_id" binding:"required"`
	Qty             float64       `json:"qty" form:"qty" default:"0"`
	IsActive        bool          `json:"is_active" form:"is_active" default:"true"`
	CreatedAt       time.Time     `json:"created_at" default:"now"`
	UpdatedAt       time.Time     `json:"updated_at" default:"now"`
	Ledger          Ledger        `gorm:"foreignKey:LedgerID;references:ID" json:"ledger"`
	ReceiveDetail   ReceiveDetail `gorm:"foreignKey:ReceiveDetailID;references:ID" json:"receive_detail"`
	Shelve          Shelve        `gorm:"foreignKey:ShelveID;references:ID" json:"shelve"`
}

func (u *Carton) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}

type CartonHistory struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	CartonID    string    `gorm:"size:21" json:"carton_id" form:"carton_id"`
	PalletNo    string    `gorm:"size:15;" json:"pallet_no" form:"pallet_no"`
	ShelveID    string    `gorm:"size:15" json:"shelve_id" form:"shelve_id"`
	Qty         float64   `json:"qty" form:"qty" default:"0"`
	IpAddress   net.IP    `gorm:"type:inet;size:15;" json:"ip_address" form:"ip_address"`
	EmpID       string    `gorm:"size:10;" json:"emp_id" form:"emp_id" binding:"required"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `json:"is_active" form:"is_active" default:"true"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (u *CartonHistory) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}
