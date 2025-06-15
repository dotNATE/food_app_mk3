package models

import "time"

type Rating struct {
	ID       int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	VendorId int64  `json:"vendor_id" gorm:"not null"`
	Score    int64  `json:"score" gorm:"default:0"`
	Review   string `json:"review"`

	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}
