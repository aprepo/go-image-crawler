package filedb

import (
	"database/sql"
	"log"
)

type FileDatabase struct {
	db *sql.DB
}

func (self FileDatabase) SetupDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./files.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS file_hashes (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"filename" TEXT NOT NULL,
		"md5_hash" TEXT NOT NULL
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (f FileDatabase) Close() {
	if f.db != nil {
		f.db.Close()
	}
}
