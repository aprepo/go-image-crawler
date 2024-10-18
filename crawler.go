package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func handleDirectory(path string, info os.FileInfo) {
	log.Println("Handling directory : " + path)
}

func handleFile(path string, info os.FileInfo) {
	log.Print("Handling file : " + path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	// Get the resulting MD5 sum
	hashInBytes := hash.Sum(nil)[:16]

	// Convert it to a hexadecimal string
	hashString := hex.EncodeToString(hashInBytes)
	log.Println(" MD5: " + hashString)
}

func callback(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Fatal(err)
		return err
	}
	if info.IsDir() {
		handleDirectory(path, info)
	} else {
		handleFile(path, info)
	}
	// fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
	return nil
}

func main() {
	args := os.Args
	path := args[1]
	fmt.Println("Search files from: " + path)
	err := filepath.Walk(path, callback)
	if err != nil {
		log.Fatal(err)
	}
}
