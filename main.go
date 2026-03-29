package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func moveFile(src, dst string) {
	os.MkdirAll(dst, os.ModePerm)
	newPath := filepath.Join(dst, filepath.Base(src))
	err := os.Rename(src, newPath)
	if err != nil {
		fmt.Println("Error moving file:", err)
	}
}

func getCategory(ext string) string {
	switch ext {
	case ".jpg", ".png", ".jpeg":
		return "Images"
	case ".pdf", ".docx", ".txt":
		return "Documents"
	case ".mp4", ".mkv":
		return "Videos"
	default:
		return "Others"
	}
}

func main() {
	var path string

	fmt.Println("==== FILE ORGANIZER CLI ====")
	fmt.Print("Enter folder path: ")
	fmt.Scanln(&path)

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Invalid path")
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fullPath := filepath.Join(path, file.Name())
		ext := filepath.Ext(file.Name())
		category := getCategory(ext)

		destFolder := filepath.Join(path, category)
		moveFile(fullPath, destFolder)
	}

	fmt.Println("✅ Files organized successfully!")
}