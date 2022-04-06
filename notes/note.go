package notes

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func SaveToDisk(n *Note, folder string) error {
	filename := filepath.Join(folder, n.Title) //title should be sanitized
	return os.WriteFile(filename, n.Body, 0600)
}

func LoadFromDisk(keyword string, folder string) (*Note, error) {
	filename, err := searchKeywordInFilename(folder, keyword)
	if err != nil {
		return nil, err
	}
	body, err := os.ReadFile(filepath.Join(folder, filename))
	if err != nil {
		return nil, err
	}
	return &Note{Title: filename, Body: body}, nil
}

// Scan a directory and if a file name contain a substring, return the first one
func searchKeywordInFilename(folder string, keyword string) (string, error) {
	items, _ := ioutil.ReadDir(folder)
	for _, item := range items {

		// Read the whole file at once
		// this is the most ineficient seach engine in the world
		// good enough for an example
		b, err := ioutil.ReadFile(filepath.Join(folder, item.Name()))
		if err != nil {
			// This is not normal but we can safely ignore it
			log.Printf("Could not read %v", item.Name())
		}
		s := string(b)

		if strings.Contains(s, keyword) {
			return item.Name(), nil
		}
	}
	return "", errors.New("no file contains this keyword")
}
