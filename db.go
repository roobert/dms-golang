package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func createDB(name string) {
	var err error
	db, err = sql.Open("sqlite3", name)
	checkErr(err)
}

func createTable(name string, fields string) {
	stmt, _ := db.Prepare("CREATE TABLE IF NOT EXISTS ? (?)")
	_, err := stmt.Exec(name, fields)
	checkErr(err)
}

func rowExists(c Client) bool {
	var exists bool

	query := fmt.Sprintf("SELECT exists (SELECT * FROM data WHERE 'name' == '%s' and 'date' == '%s')", c.Name, c.Date)
	err := db.QueryRow(query).Scan(&exists)
	checkErr(err)

	return exists
}

func checkForMultipleRecords(c Client) error {
	var count int

	// date.Date()
	query := fmt.Sprintf("SELECT count(*) FROM data WHERE 'name' == '%s' and 'date' == date(%s))", c.Name, c.Date)
	rows := db.QueryRow(query)
	err := rows.Scan(&count)
	checkErr(err)

	if count > 1 {
		return nil
	} else {
		return fmt.Errorf("error: found more than one row for client with date: %s/%s: %i", c.Name, c.TimeStamp, count)
	}
}
