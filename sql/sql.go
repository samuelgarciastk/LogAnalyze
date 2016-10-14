package sql

import (
	. "LogAnalyze/common"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

//Open is used to open the database
func Open() {
	var err error
	db, err = sql.Open("sqlite3", "data.db")
	CheckErr(err, "Open database failed!")
}

//Close is used to close the database
func Close() {
	db.Close()
}

//Insert is used to insert data slice into the database
//'query' is the query sequence
func InsertSlice(data []string, query string) {
	tx, err := db.Begin()
	CheckErr(err)
	stmt, err := tx.Prepare(query)
	CheckErr(err)
	defer stmt.Close()
	for _, value := range data {
		stmt.Exec(value)
	}
	tx.Commit()
}

//Insert is used to insert data map into the database
//'query' is the query sequence
func InsertMap(data map[string]interface{}, query string) {
	tx, err := db.Begin()
	CheckErr(err)
	stmt, err := tx.Prepare(query)
	CheckErr(err)
	defer stmt.Close()
	for k := range data {
		stmt.Exec(k)
	}
	tx.Commit()
}

//Exec is used to execute queries
func Exec(queries ...string) {
	for _, v := range queries {
		_, err := db.Exec(v)
		CheckErr(err)
	}
}
