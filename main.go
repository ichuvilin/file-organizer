package main

import "fmt"

var DefaultRules = map[string]string{
	".jpg":  "Images",
	".jpeg": "Images",
	".pdf":  "Documents",
	".mp3":  "Music",
}

func main() {
	folder := DefaultRules[".jpg"]
	fmt.Println(folder)
}
