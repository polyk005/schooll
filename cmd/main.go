package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	OnlineSchooleProject "github.com/polyk005/school"
	"github.com/polyk005/school/pkg/create"
	"github.com/polyk005/school/pkg/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	dbConn, err := db.NewDB()
	if err != nil {
		logrus.Fatalf("error connecting to database: %s", err.Error())
	}
	defer dbConn.Close()

	if err := createTables(dbConn.GetDB()); err != nil {
		logrus.Fatalf("error creating tables: %s", err.Error())
	}

	logrus.Info("Starting server...")

	srv := new(OnlineSchooleProject.Server)
	if err := srv.Run(viper.GetString("port")); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func createTables(db *sqlx.DB) error {
	tables := []create.TableDefinition{
		{
			Name: "users",
			Columns: map[string]string{
				"id":         "SERIAL PRIMARY KEY",
				"username":   "VARCHAR(100) NOT NULL",
				"email":      "VARCHAR(100) NOT NULL UNIQUE",
				"created_at": "TIMESTAMP DEFAULT CURRENT_TIMESTAMP",
			},
		},
		{
			Name: "courses",
			Columns: map[string]string{
				"id":          "SERIAL PRIMARY KEY",
				"title":       "VARCHAR(255) NOT NULL",
				"description": "TEXT",
				"created_at":  "TIMESTAMP DEFAULT CURRENT_TIMESTAMP",
			},
		},
	}

	for _, table := range tables {
		if err := create.CreateTable(db, table); err != nil {
			return fmt.Errorf("error creating table %s: %w", table.Name, err)
		}
	}

	return nil
}
