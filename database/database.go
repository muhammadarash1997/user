package database

import (
	"database/sql"
	"fmt"
)

func StartConnection() (*sql.DB, error) {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", "root", "root", "db:3306", "user_management")

	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

//create migrate file
//migrate create -ext sql -dir db/migrations create_table_users

//migrate up:
//migrate -database "mysql://root@tcp(localhost:3306)/user" -path db/migrations up