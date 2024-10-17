package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func callback(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
	return nil
}

func main() {
	args := os.Args
	path := args[1]
	fmt.Println("Search files from: " + path)
	err := filepath.Walk(path, callback)
	if err != nil {
		fmt.Println(err)
	}
}
