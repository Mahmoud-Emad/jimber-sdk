package server

import (
	"fmt"

	"gorm.io/gorm"
	// "gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var ConnStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
