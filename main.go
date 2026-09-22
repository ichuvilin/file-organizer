package main

import (
	"errors"
	"fmt"
	"os"
)

var DefaultRules = map[string]string{
	".jpg":  "Images",
	".jpeg": "Images",
	".pdf":  "Documents",
	".mp3":  "Music",
}

type FileOrganizer struct {
	sourceDir      string
	rulesMap       map[string]string
	processedFiles int
	logFile        *os.File
}

func NewFileOrganizer(sourceDir string) (*FileOrganizer, error) {
	if sourceDir == "" {
		return nil, errors.New("SourceDir not recognize")
	}
	info, err := os.Stat(sourceDir)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, errors.New("sourceDir is not a dir")
	}

	return &FileOrganizer{sourceDir: sourceDir, rulesMap: DefaultRules}, nil
}

func main() {
	folder := DefaultRules[".jpg"]
	fmt.Println(folder)
}
