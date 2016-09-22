package sql

import (
	. "LogAnalyze/common"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Insert(data []string) {
	db, err := sql.Open("sqlite3", "template.db")
	CheckErr(err, "Open database failed!")
	defer db.Close()

	_, err = db.Exec("DELETE FROM template")
	CheckErr(err)
	_, err = db.Exec("UPDATE sqlite_sequence SET seq=0 WHERE name='template'")
	CheckErr(err)

	tx, err := db.Begin()
	CheckErr(err)
	stmt, err := tx.Prepare("INSERT INTO template(event) values(?)")
	CheckErr(err)
	defer stmt.Close()
	for _, value := range data {
		_, err = stmt.Exec(value)
		LogErr(err, value)
	}
	tx.Commit()
}
