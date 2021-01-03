package short

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// fileRepo is a repository that stores Link in a file
type fileRepo struct {
	db   string
	file *os.File
}

// NewFileRepo makes a new file repository
func NewFileRepo(db string) *Repo {
	var repo Repo
	repo = &fileRepo{
		db:   db,
		file: nil,
	}
	return &repo
}

// Init intialize FileRepo
func (f *fileRepo) Init() error {
	file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	f.file = file
	return nil
}

// Save saves a short link to a file
func (f *fileRepo) Save(l *Link) error {
	if isValid, _ := regexp.Match("([^\\|\\[\\]])", []byte(l.Key)); !isValid {
		return fmt.Errorf("Invalid Key")
	}
	if isValid, _ := regexp.Match("([^\\|\\[\\]])", []byte(l.Val)); !isValid {
		return fmt.Errorf("Invalid Value")
	}
	_, err := f.file.Write([]byte(l.String() + "\n"))
	if err != nil {
		return err
	}
	return nil
}

// Get searches for a key in database and returns the value
func (f *fileRepo) Get(key string) (string, error) {
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
