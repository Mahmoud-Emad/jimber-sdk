// Package models for database models
package models

import (
	"errors"
	"fmt"

	"github.com/Mahmoud-Emad/envserver/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB struct hold db instance
type Database struct {
	db *gorm.DB
}

// NewDatabase create and return new Database struct.
func NewDatabase() Database {
	return Database{}
}

// Connect connects to database server.
func (d *Database) Connect(dbConfig internal.DatabaseConfiguration) error {
	if dbConfig.Host == "" || dbConfig.Port == 0 || dbConfig.User == "" || dbConfig.Password == "" || dbConfig.Name == "" {
		return errors.New("invalid database configuration")
	}

	ConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)

	gormDB, err := gorm.Open(postgres.Open(ConnStr), &gorm.Config{})

	if err != nil {
		return err
	}

	d.db = gormDB
	return nil
}

// Migrate migrates db schema
func (d *Database) Migrate() error {
	err := d.db.AutoMigrate(&User{}, &Project{}, &EnvironmentKey{})
	if err != nil {
		return err
	}
	return nil
}

// Create new user object inside the daabase.
func (d *Database) CreateUser(u *User) error {
	result := d.db.Create(&u)
	return result.Error
}

// GetUserByEmail returns user by its email
func (d *Database) GetUserByEmail(email string) (User, error) {
	var u User
	query := d.db.First(&u, "email = ?", email)
	return u, query.Error
}

// DeleteUserByEmail deletes a user by their email
func (d *Database) DeleteUserByEmail(email string) error {
	result := d.db.Unscoped().Where("email = ?", email).Delete(&User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Create new project object inside the daabase.
func (d *Database) CreateProject(p *Project) error {
	result := d.db.Create(&p)
	return result.Error
}

// DeleteProjectByName deletes a project by it's name
func (d *Database) DeleteProjectByName(name string) error {
	result := d.db.Unscoped().Where("name = ?", name).Delete(&Project{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetProjectByName returns user by its name
func (d *Database) GetProjectByName(name string) (Project, error) {
	var p Project
	query := d.db.First(&p, "name = ?", name)
	return p, query.Error
}

// Create new EnvironmentKey object inside the daabase.
func (d *Database) CreateEnvKey(env *EnvironmentKey) error {
	result := d.db.Create(&env)
	return result.Error
}

// DeleteEnvironmentKeyByKeyName Delete an EnvironmentKey by it's key name.
func (d *Database) DeleteEnvKeyByKeyName(keyName string) error {
	result := d.db.Unscoped().Where("key = ?", keyName).Delete(&EnvironmentKey{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetEnvKeyByKeyName returns user by its key name
func (d *Database) GetEnvKeyByKeyName(keyName string) (EnvironmentKey, error) {
	var env EnvironmentKey
	query := d.db.First(&env, "key = ?", keyName)
	return env, query.Error
}