package short

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// FileRepo is a repository that stores Link in a file
type FileRepo struct {
	db   string
	file *os.File
}

// Init intialize FileRepo
func (f *FileRepo) Init() error {
	file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	f.file = file
	return nil
}

// Save saves a short link to a file
func (f *FileRepo) Save(l *Link) error {
	if isValid, _ := regexp.Match("([^\\|\\[\\]])", []byte(l.key)); !isValid {
		return fmt.Errorf("Invalid Key")
	}
	if isValid, _ := regexp.Match("([^\\|\\[\\]])", []byte(l.val)); !isValid {
		return fmt.Errorf("Invalid Value")
	}
	_, err := f.file.Write([]byte(l.String() + "\n"))
	if err != nil {
		return err
	}
	return nil
}

// Get searches for a key in database and returns the value
func (f *FileRepo) Get(key string) (string, error) {
	scanner := bufio.NewScanner(f.file)
	for scanner.Scan() {
		line := scanner.Text()
		strings := strings.Split(line, " | ")
		if key == strings[0] {
			return strings[1], nil
		}
	}
	return "", fmt.Errorf("No value found")
}
