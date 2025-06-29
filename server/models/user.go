package models

import "time"

type User struct {
	ID    int64  `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"unique;not null"`
	Email string `gorm:"unique;not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
