package repositories

import (
	"time"

	"example.com/hasher/internal/domain"
	"github.com/google/uuid"
)

// Hasher repository
type Hasher struct{}

// Instantiate hasher repository
func NewHasher() *Hasher {
	return &Hasher{}
}

// Generate new hash
func (h *Hasher) Generate() *domain.Hash {
	return &domain.Hash{
		Value:     uuid.New(),
		CreatedAt: time.Now(),
	}
}
