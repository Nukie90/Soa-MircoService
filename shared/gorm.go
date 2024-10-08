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

func (d *Database) PostgresConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgresql://virtual_banking_db_owner:%s@%s.ap-southeast-1.aws.neon.tech/%s?sslmode=require", d.Password, d.ComputeID, d.DBName)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}

	err = conn.AutoMigrate(&entity.User{}, &entity.Account{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
		return nil, err
	}

	return conn, nil
}
