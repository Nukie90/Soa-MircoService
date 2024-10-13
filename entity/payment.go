package entity

import (
	cryptoRand "crypto/rand"
	"math/big"
	"time"

	"github.com/oklog/ulid/v2"

	"gorm.io/gorm"
)

type Payment struct {
	ID              string    `gorm:"primaryKey"`
	SourceAccountID string    `gorm:"not null"`
	SourceAccount   Account   `gorm:"foreignKey:SourceAccountID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReferenceCode   string    `gorm:"not null"`
	Amount          float64   `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt
}

func (Payment) TableName() string {
	return "payments"
}

func (c *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(cryptoRand.Reader, 0)
	ulid := ulid.MustNew(ulid.Timestamp(t), entropy)

	// Convert the ULID to a big.Int
	bigInt := new(big.Int).SetBytes(ulid[:])

	// Get the last 10 digits
	modulus := new(big.Int).SetInt64(10000000000)
	result := new(big.Int).Mod(bigInt, modulus)

	// Convert the result to a string
	id := result.String()

	// Ensure the result is 10 digits by padding with leading zeros if necessary
	for len(id) < 10 {
		id = "0" + id
	}

	c.ID = id
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	return
}

func (c *Payment) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	return
}