package short

import (
	"fmt"
)

// Link defines a key value pair
type Link struct {
	key string
	val string
}

func (s *Link) String() string {
	return fmt.Sprintf("%v | %v", s.key, s.val)
}

// Repo defines an interface that all Link Repository need to have
type Repo interface {
	Init() error
	Get(key string) (string, error)
	Save(*Link) error
}
