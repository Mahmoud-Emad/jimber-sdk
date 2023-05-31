package server

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (stg *Storage) Connect() *gorm.DB {
	ConnStr := ConnStr
	db, err := gorm.Open(postgres.Open(ConnStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("|-| Failed to connect to the database: %v", err)
	}
	return db
}

func NewStorage() *Storage {
	return &Storage{
		DB: &gorm.DB{},
	}
}
