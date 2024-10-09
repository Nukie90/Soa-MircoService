package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Address   string    `gorm:"not null"`
	BirthDate time.Time `gorm:"not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
