package metadata

import (
	"fmt"
	"time"
)

type FileMetadata struct {
	FilePath  string
	Filename  string
	FileType  string
	MD5hash   string
	TimeTaken time.Time
}

func (f FileMetadata) DisplaySummary() {
	fmt.Println("FilePath: " + f.FilePath)
	fmt.Println("MD5: " + f.MD5hash)
}
