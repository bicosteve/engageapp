package utils

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Contains DB connection helpers

func ConnectDB(dbString string) *sql.DB {
	Log("INFO", "dbs", "connecting to db...")
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		Log("ERROR", "dbs", "unable to connect to db because of %v", err)
		os.Exit(1)
	}

	if err = db.Ping(); err != nil {
		Log("ERROR", "dbs", "unable to ping db because of %v", err)
		os.Exit(1)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(time.Second * 20)
	db.SetConnMaxLifetime(time.Minute)

	Log("INFO", "dbs", "connected to db successfully")

	return db
}
