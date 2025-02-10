package create

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type TableDefinition struct {
	Name    string
	Columns map[string]string
}

func CreateTable(db *sqlx.DB, table TableDefinition) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", table.Name)

	i := 0
	for colName, colType := range table.Columns {
		if i > 0 {
			query += ", "
		}
		query += fmt.Sprintf("%s %s", colName, colType)
		i++
	}
	query += ");"

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", table.Name, err)
	}

	log.Printf("Table %s created successfully\n", table.Name)
	return nil
}
