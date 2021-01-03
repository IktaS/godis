package short

import (
	"fmt"
)

// Link defines a key value pair
type Link struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func (s *Link) String() string {
	return fmt.Sprintf("%v | %v", s.Key, s.Val)
}

// Repo defines an interface that all Link Repository need to have
type Repo interface {
	Init() error
	Get(key string) (string, error)
	Save(*Link) error
}
