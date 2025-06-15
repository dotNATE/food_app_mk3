package models

import "time"

type AuthIdentity struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	UserId   int64  `gorm:"not null;index"`
	Password string `gorm:"not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserId"`
}
