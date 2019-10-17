package database

import (
	"database/sql"
	"github.com/BRUHItsABunny/bunnlog"
	_ "github.com/go-sql-driver/mysql"
)

type ExampleDatabase struct {
	DB  *sql.DB
	Log *bunnlog.BunnyLog
}

func GetExampleDatabase(bLog *bunnlog.BunnyLog) (*ExampleDatabase, error) {

	dbUsername := ""
	dbPassword := ""
	dbAddress := ""
	dbDatabase := ""

	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbAddress+":3306)/"+dbDatabase)
	if err == nil {
		return nil, err
	} else {
		return &ExampleDatabase{DB: db, Log: bLog}, nil
	}
}
