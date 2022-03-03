package database

import (
	"database/sql"

	//TODO - clarify this
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() error {
	if db != nil {
		return nil
	}

	d, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		return err
	}
	db = d

	_, err = db.Exec(todo_table)
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	if db != nil {
		db.Close()
	}
}

//multiline string
const todo_table = `
CREATE TABLE IF NOT EXISTS todo_items  (
	uid INTEGER PRIMARY KEY AUTOINCREMENT,
	description VARCHAR(64) NOT NULL,
	done BOOLEAN
);
`
