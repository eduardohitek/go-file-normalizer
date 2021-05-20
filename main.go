package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {
	// var files []string

	root := "/Users/hitek/dev/chars"
	file, err := os.Open(root)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, err := file.Readdirnames(0) // 0 to read all files and folders
	if err != nil {
		log.Fatalf("failed reading directory: %s", err)
	}
	for _, name := range list {
		oldName := name
		fmt.Println("Old Name - ", oldName)
		newName := normalizeFileName(oldName)
		fmt.Println("New Name - ", newName)
		if err != nil {
			log.Printf("error renaming file: %s", err)
			continue
		}
		renameFile(root, oldName, newName)
		fmt.Println("File names have been changed")
	}
}

func renameFile(root string, src string, dest string) {
	os.Rename(filepath.Join(root, src), filepath.Join(root, dest))
}

func normalizeFileName(fileName string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	fileNameNormalized, _, _ := transform.String(t, fileName)
	fileNameNormalized = strings.ReplaceAll(fileNameNormalized, " ", "_")
	return fileNameNormalized
}
