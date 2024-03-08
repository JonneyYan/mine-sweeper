package storage

import (
	ms "infinitemining.com/minesweeper"
)

type Storage interface {
	Total() int
	Save(item interface{})
	Load(index int) (interface{}, error)
	LoadLatest() (interface{}, error)
	LoadAll() []interface{}
}

type MemStorage struct {
	Store []interface{}

	Limit int
}

func NewMemStorage(limit int) *MemStorage {
	s := &MemStorage{
		Limit: limit,
	}

	return s
}

func (s *MemStorage) Total() int {
	return len(s.Store)
}
func (s *MemStorage) Save(item interface{}) {
	s.Store = append(s.Store, item)
	copy(s.Store[1:], s.Store)
	s.Store[0] = item

	if s.Limit < len(s.Store) {
		s.Store = s.Store[0 : len(s.Store)-1]
	}

}
func (s *MemStorage) Load(index int) (interface{}, error) {
	if index > len(s.Store) {
		return nil, ms.ErrOverflow
	}
	return s.Store[index], nil
}
func (s *MemStorage) LoadLatest() (interface{}, error) {
	if len(s.Store) == 0 {
		return nil, ms.ErrEmpty
	}
	return s.Store[0], nil
}

func (s *MemStorage) LoadAll() []interface{} {
	return s.Store
}
