package tests

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/polyk005/school/pkg/create"
	"github.com/polyk005/school/pkg/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTables(t *testing.T) {
	config := map[string]string{
		"db.username": "postgres",
		"db.host":     "localhost",
		"db.port":     "5432",
		"db.dbname":   "test_db",
		"db.sslmode":  "disable",
		"db.password": "-",
	}

	// Подключение к базе данных
	dbConn, err := db.NewDBWithConfig(config)
	require.NoError(t, err, "Failed to connect to database")
	defer dbConn.Close()

	// Создание таблиц
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
		err := create.CreateTable(dbConn.GetDB(), table)
		assert.NoError(t, err, "Failed to create table %s", table.Name)
	}

	// Проверка, что таблицы созданы
	var exists bool
	err = dbConn.GetDB().Get(&exists, "SELECT EXISTS (SELECT FROM pg_tables WHERE tablename = 'users')")
	assert.NoError(t, err, "Failed to check if table 'users' exists")
	assert.True(t, exists, "Table 'users' should exist")

	err = dbConn.GetDB().Get(&exists, "SELECT EXISTS (SELECT FROM pg_tables WHERE tablename = 'courses')")
	assert.NoError(t, err, "Failed to check if table 'courses' exists")
	assert.True(t, exists, "Table 'courses' should exist")
}
