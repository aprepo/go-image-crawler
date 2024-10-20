package main

import (
	"crawler/filedb"
	"crawler/filescanner"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	args := os.Args
	path := args[1]
	fmt.Println("Search files from: " + path)

	fileDb := filedb.FileDatabase{}
	fileDb.SetupDatabase()
	defer fileDb.Close()

	scanner := filescanner.FileScanner{
		Db: &fileDb,
	}
	scanner.ScanFiles(path)
}
