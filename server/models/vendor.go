package models

import "time"

type Vendor struct {
	ID            int64   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string  `json:"name" gorm:"not null"`
	Description   string  `json:"description"`
	AverageRating float64 `json:"average_rating" gorm:"type:decimal(3,2);not null;default:0"`

	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`

	Ratings []Rating `json:"ratings,omitempty" gorm:"foreignKey:VendorId"`
}
