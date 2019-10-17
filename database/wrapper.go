package database

import "database/sql"

func (db *ExampleDatabase) ExampleSelectQuery(userId string) error {
	var err error
	var row *sql.Row

	q := "SELECT ID FROM users WHERE ID=?"
	row = db.DB.QueryRow(q, userId)
	err = row.Scan()

	return err
}
