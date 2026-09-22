package main

import (
	"errors"
	"fmt"
	"log"
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

func (fo *FileOrganizer) initLog() error {
	fileName := "organizer.log"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	log.SetOutput(file)
	fo.logFile = file
	return nil
}

func (fo *FileOrganizer) Close() error {
	if fo.logFile == nil {
		return nil
	}

	if err := fo.logFile.Close(); err != nil {
		return err
	}

	return nil
}

func (fo *FileOrganizer) logSuccess(message string) {
	log.Printf("[SUCCESS] %s\n", message)
}

func (fo *FileOrganizer) logError(message string) {
	log.Printf("[ERROR] %s\n", message)
}

func main() {
	folder := DefaultRules[".jpg"]
	fmt.Println(folder)
}
