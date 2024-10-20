package main

import (
	"crawler/filedb"
	"crawler/handlers"
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

type FileScanner struct {
	db *filedb.FileDatabase
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

	fileDb := filedb.FileDatabase{}
	fileDb.SetupDatabase()
	defer fileDb.Close()

	scanner := FileScanner{
		db: &fileDb,
	}
	scanner.scanFiles(path)
}
