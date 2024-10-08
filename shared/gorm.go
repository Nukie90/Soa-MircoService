package shared

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"microservice/db/entity"
)

type Database struct {
	ComputeID string
	Password  string
	DBName    string
}

func NewDatabase(computeID, password, dbName string) *Database {
	return &Database{
		ComputeID: computeID,
		Password:  password,
		DBName:    dbName,
	}
}

func (d *Database) PostgresConnection(db *gorm.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgresql://virtual_banking_db_owner:%s@%s.ap-southeast-1.aws.neon.tech/%s?sslmode=require", d.ComputeID, d.Password, d.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("Connected to database")

	err = db.AutoMigrate(&entity.Account{}, &entity.User{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
