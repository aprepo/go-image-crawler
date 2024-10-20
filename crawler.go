package main

import (
	"crawler/handlers"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func callback(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Fatal(err)
		return err
	}
	if info.IsDir() {
		handlers.HandleDirectory(path, info)
	} else {
		meta := handlers.HandleFile(path, info)
		meta.DisplaySummary()
		// TODO: Store to database
	}
	return nil
}

func setupDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./files.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type FileScanner struct {
	db *sql.DB
}

func (self FileScanner) scanFiles(path string) {
	err := filepath.Walk(path, callback)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	args := os.Args
	path := args[1]
	fmt.Println("Search files from: " + path)

	db := setupDatabase()
	defer db.Close()

	scanner := FileScanner{
		db: db,
	}
	scanner.scanFiles(path)
}
