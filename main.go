package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
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

func (fo *FileOrganizer) moveFile(sourcePath, targetDir string) error {
	fullPath := filepath.Join(fo.sourceDir, targetDir)
	fileName := filepath.Base(sourcePath)

	err := os.MkdirAll(fullPath, 0644)
	if err != nil {
		return err
	}

	_, err = os.Stat(filepath.Join(fullPath, fileName))
	if err == nil {
		ext := filepath.Ext(fileName)
		name := strings.TrimSuffix(fileName, ext)
		fileName = fmt.Sprintf("%s_%s%s", name, time.Now().Format("2006-01-02_15-04-05"), ext)
	}
	err = os.Rename(sourcePath, filepath.Join(fullPath, fileName))
	if err != nil {
		return err
	}
	return nil
}

func (fo *FileOrganizer) Organize() error {
	err := fo.initLog()
	if err != nil {
		return err
	}
	return filepath.WalkDir(fo.sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		if filepath.Dir(path) != fo.sourceDir {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(d.Name()))
		folder, ok := fo.rulesMap[ext]
		if ok {
			fo.moveFile(path, folder)
			fo.processedFiles += 1
		}

		return nil
	})
}

func main() {
	folder := DefaultRules[".jpg"]
	fmt.Println(folder)
}
