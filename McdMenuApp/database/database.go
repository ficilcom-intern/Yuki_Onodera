package database

import (
	"database/sql"
)

type Item struct {
	Kind          string
	Name          string
	Energy        float64
	Protein       float64
	Fat           float64
	Carbohydrates float64
}

func createUsersTable(db *sql.DB) error {
	query := `
        CREATE TABLE IF NOT EXISTS Item (
            kind          TEXT,
			name          TEXT,
			energy        FLOAT,
			protein       FLOAT,
			fat           FLOAT,
			carbohydrates FLOAT
        )
    `
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
