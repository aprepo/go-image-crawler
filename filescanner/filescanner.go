package filescanner

import (
	"crawler/filedb"
	"crawler/handlers"
	"log"
	"os"
	"path/filepath"
)

type FileScanner struct {
	Db *filedb.FileDatabase
}

func (self FileScanner) ScanFiles(path string) {
	err := filepath.Walk(path, self.callback)
	if err != nil {
		log.Fatal(err)
	}
}

func (fs FileScanner) callback(path string, info os.FileInfo, err error) error {
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
