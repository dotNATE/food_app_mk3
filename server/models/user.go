package models

import "time"

type User struct {
	ID    int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"unique;not null"`
	Email string `json:"email" gorm:"unique;not null"`

	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}
