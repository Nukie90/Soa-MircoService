package entity

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type User struct {
	Id        ulid.ULID `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Address   string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	BirthDate time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

func (User) TableName() string {
	return "clients"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	u.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
