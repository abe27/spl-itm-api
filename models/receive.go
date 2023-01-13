package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Receive struct {
	ID              string          `gorm:"primaryKey;size:21" json:"id"`
	DownloadID      *string         `gorm:"not null" form:"download_id" json:"download_id"`
	ReceiveTypeID   *string         `gorm:"not null" form:"receive_type_id" json:"receive_type_id"`
	ReceiveDate     time.Time       `gorm:"type:date" json:"receive_date" form:"receive_date" binding:"required"`
	TransferOutNo   string          `gorm:"not null;unique;size:15" json:"transfer_out_no" form:"transfer_out_no" binding:"required"`
	TexNo           string          `gorm:"size:15;" json:"tex_no" form:"tex_no"`
	Item            int64           `json:"item" form:"item" default:"0"`
	PlanCtn         int64           `json:"plan_ctn" form:"plan_ctn" default:"0"`
	ReceiveCtn      int64           `json:"receive_ctn" form:"receive_ctn" default:"0"`
	IsSync          bool            `json:"is_sync" form:"is_sync" default:"true"`
	IsActive        bool            `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt       time.Time       `json:"created_at" default:"now"`
	UpdatedAt       time.Time       `json:"updated_at" default:"now"`
	DownloadMailBox DownloadMailBox `gorm:"foreignKey:DownloadID;references:ID" json:"download"`
	ReceiveType     ReceiveType     `gorm:"foreignKey:ReceiveTypeID;references:ID" json:"receive_type"`
	ReceiveDetail   []ReceiveDetail `json:"receive_detail"`
}

func (u *Receive) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}

type ReceiveDetail struct {
	ID               string             `gorm:"primaryKey;size:21" json:"id"`
	ReceiveID        *string            `gorm:"not null;" form:"receive_id" json:"receive_id"`
	LedgerID         *string            `gorm:"not null;" form:"ledger_id" json:"ledger_id"`
	PlanQty          int64              `json:"plan_qty" form:"plan_qty"`
	PlanCtn          int64              `json:"plan_ctn" form:"plan_ctn"`
	IsActive         bool               `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt        time.Time          `json:"created_at" default:"now"`
	UpdatedAt        time.Time          `json:"updated_at" default:"now"`
	Receive          Receive            `gorm:"foreignKey:ReceiveID;references:ID" json:"receive"`
	Ledger           Ledger             `gorm:"foreignKey:LedgerID;references:ID" json:"ledger"`
	CartonNotReceive []CartonNotReceive `json:"receive_carton"`
	// Cartons   []Carton  `json:"carton"`
}

func (u *ReceiveDetail) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}

type CartonNotReceive struct {
	ID              string        `gorm:"primaryKey;size:21" json:"id,omitempty"`
	ReceiveDetailID string        `gorm:"not null;" json:"receive_detail_id,omitempty" form:"receive_detail_id" binding:"required"`
	TransferOutNo   string        `gorm:"not null;size:25" json:"transfer_out_no,omitempty" form:"transfer_out_no" binding:"required"`
	PartNo          string        `gorm:"not null;" json:"part_no,omitempty" form:"part_no" binding:"required"`
	LotNo           string        `gorm:"not null;size:8;" json:"lot_no,omitempty" form:"lot_no" binding:"required"`
	SerialNo        string        `gorm:"not null;size:10;" json:"serial_no,omitempty" form:"serial_no" binding:"required"`
	Qty             int64         `json:"qty,omitempty" form:"qty" binding:"required"`
	IsReceived      bool          `json:"is_received,omitempty" form:"is_received"`
	IsSync          bool          `json:"is_sync,omitempty" form:"is_sync" default:"false"`
	CreatedAt       time.Time     `json:"created_at,omitempty" default:"now"`
	UpdatedAt       time.Time     `json:"updated_at,omitempty" default:"now"`
	ReceiveDetail   ReceiveDetail `gorm:"foreignKey:ReceiveDetailID;references:ID" json:"receive_detail"`
}

func (u *CartonNotReceive) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}
