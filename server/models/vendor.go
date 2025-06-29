package models

import "time"

type Vendor struct {
	ID            int64  `gorm:"primaryKey;autoIncrement"`
	Name          string `gorm:"not null"`
	Description   string
	AverageRating float64 `gorm:"type:decimal(4,2);not null;default:0"`

	CreatedBy int64     `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Ratings       []Rating `gorm:"foreignKey:VendorId"`
	CreatedByUser User     `gorm:"foreignKey:CreatedBy;references:ID"`
}
