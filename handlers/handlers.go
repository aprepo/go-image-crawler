package handlers

import (
	"crawler/metadata"
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func HandleDirectory(path string, info os.FileInfo) {
	log.Println("Handling directory : " + path)
}

func HandleFile(path string, info os.FileInfo) metadata.FileMetadata {
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

	// Create a FileMetadata object
	meta := metadata.FileMetadata{
		FilePath: path,
		MD5hash:  hashString,
	}

	return meta
}
