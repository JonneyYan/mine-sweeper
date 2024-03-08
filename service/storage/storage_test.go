package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {

	s := NewMemStorage(5)

	s.Save(1)
	s.Save(2)
	s.Save(3)
	s.Save(4)
	s.Save(5)

	v0, _ := s.Load(0)
	assert.Equal(t, 5, v0)

	s.Save(6)
	v4, _ := s.Load(4)
	assert.Equal(t, 2, v4)

	total := s.Total()
	assert.Equal(t, 5, total)

	v0, _ = s.LoadLatest()
	assert.Equal(t, 6, v0)

	res := s.LoadAll()
	assert.Equal(t, 5, len(res))

}
