package set

import (
	"strconv"
	"strings"
)

// Set represents a set data structure
type Set struct {
	Elements map[int]struct{}
}

// NewSet returns a new set
func NewSet() *Set {
	return &Set{Elements: make(map[int]struct{})}
}

// Add adds an element to the set
func (s *Set) Add(element int) {
	s.Elements[element] = struct{}{}
}

// Contains returns true if the set contains the element
func (s *Set) Contains(element int) bool {
	_, exists := s.Elements[element]
	return exists
}

// List prints the elements of the set
func (s *Set) List() string {
	keys := make([]string, 0, len(s.Elements))
	for k := range s.Elements {
		keys = append(keys, strconv.FormatInt(int64(k), 10))
	}
	return strings.Join(keys, ",")
}
