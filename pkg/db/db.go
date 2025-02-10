package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type DB struct {
	*sqlx.DB
}

// NewDB создает новое соединение с базой данных
func NewDB() (*DB, error) {
	return NewDBWithConfig(map[string]string{
		"db.username": viper.GetString("db.username"),
		"db.host":     viper.GetString("db.host"),
		"db.port":     viper.GetString("db.port"),
		"db.dbname":   viper.GetString("db.dbname"),
		"db.sslmode":  viper.GetString("db.sslmode"),
		"db.password": viper.GetString("db.password"),
	})
}

// NewDBWithConfig создает новое соединение с базой данных на основе переданной конфигурации
func NewDBWithConfig(config map[string]string) (*DB, error) {
	connStr := fmt.Sprintf(
		"user=%s host=%s port=%s dbname=%s sslmode=%s password=%s",
		config["db.username"],
		config["db.host"],
		config["db.port"],
		config["db.dbname"],
		config["db.sslmode"],
		config["db.password"],
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to the database")
	return &DB{db}, nil
}

func (d *DB) GetDB() *sqlx.DB {
	return d.DB
}
